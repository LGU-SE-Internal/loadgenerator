package stats

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

// 请求延迟统计
type LatencyStats struct {
	URL        string
	Method     string
	Latencies  []time.Duration
	Count      int
	MaxLatency time.Duration
	MinLatency time.Duration
	AvgLatency time.Duration
	mu         sync.RWMutex
}

func NewLatencyStats(url, method string) *LatencyStats {
	return &LatencyStats{
		URL:        url,
		Method:     method,
		Latencies:  make([]time.Duration, 0),
		MinLatency: time.Duration(^uint64(0) >> 1), // 最大值
	}
}

func (ls *LatencyStats) AddLatency(latency time.Duration) {
	ls.mu.Lock()
	defer ls.mu.Unlock()

	ls.Latencies = append(ls.Latencies, latency)
	ls.Count++

	if latency > ls.MaxLatency {
		ls.MaxLatency = latency
	}
	if latency < ls.MinLatency {
		ls.MinLatency = latency
	}

	// 计算平均延迟
	total := time.Duration(0)
	for _, l := range ls.Latencies {
		total += l
	}
	ls.AvgLatency = total / time.Duration(ls.Count)
}

func (ls *LatencyStats) GetPercentile(percentile float64) time.Duration {
	ls.mu.RLock()
	defer ls.mu.RUnlock()

	if len(ls.Latencies) == 0 {
		return 0
	}

	// 创建副本并排序
	sortedLatencies := make([]time.Duration, len(ls.Latencies))
	copy(sortedLatencies, ls.Latencies)
	sort.Slice(sortedLatencies, func(i, j int) bool {
		return sortedLatencies[i] < sortedLatencies[j]
	})

	index := int(float64(len(sortedLatencies)) * percentile)
	if index >= len(sortedLatencies) {
		index = len(sortedLatencies) - 1
	}

	return sortedLatencies[index]
}

func (ls *LatencyStats) GetStats() (min, max, avg time.Duration, p50, p95, p99 time.Duration) {
	ls.mu.RLock()
	defer ls.mu.RUnlock()

	if ls.Count == 0 {
		return 0, 0, 0, 0, 0, 0
	}

	return ls.MinLatency, ls.MaxLatency, ls.AvgLatency,
		ls.GetPercentile(0.5), ls.GetPercentile(0.95), ls.GetPercentile(0.99)
}

// 全局延迟统计管理器
type LatencyManager struct {
	stats map[string]*LatencyStats
	mu    sync.RWMutex
}

func NewLatencyManager() *LatencyManager {
	return &LatencyManager{
		stats: make(map[string]*LatencyStats),
	}
}

func (lm *LatencyManager) GetOrCreateStats(url, method string) *LatencyStats {
	key := fmt.Sprintf("%s:%s", method, url)

	lm.mu.RLock()
	stats, exists := lm.stats[key]
	lm.mu.RUnlock()

	if !exists {
		lm.mu.Lock()
		// 双检查锁定
		if stats, exists = lm.stats[key]; !exists {
			stats = NewLatencyStats(url, method)
			lm.stats[key] = stats
		}
		lm.mu.Unlock()
	}

	return stats
}

func (lm *LatencyManager) GetAllStats() []*LatencyStats {
	lm.mu.RLock()
	defer lm.mu.RUnlock()

	result := make([]*LatencyStats, 0, len(lm.stats))
	for _, stats := range lm.stats {
		result = append(result, stats)
	}
	return result
}

// GetTopSlowStats 获取延迟最高的前N个请求统计
func (lm *LatencyManager) GetTopSlowStats(limit int) []*LatencyStats {
	lm.mu.RLock()
	defer lm.mu.RUnlock()

	result := make([]*LatencyStats, 0, len(lm.stats))
	for _, stats := range lm.stats {
		result = append(result, stats)
	}

	// 按最大延迟排序，从高到低
	sort.Slice(result, func(i, j int) bool {
		return result[i].MaxLatency > result[j].MaxLatency
	})

	// 只返回前limit个
	if limit > 0 && len(result) > limit {
		result = result[:limit]
	}

	return result
}

func (lm *LatencyManager) HasSlowRequests(threshold time.Duration) bool {
	lm.mu.RLock()
	defer lm.mu.RUnlock()

	for _, stats := range lm.stats {
		if stats.MaxLatency > threshold {
			return true
		}
	}
	return false
}

func (lm *LatencyManager) GetSlowRequestsCount(threshold time.Duration) int {
	lm.mu.RLock()
	defer lm.mu.RUnlock()

	count := 0
	for _, stats := range lm.stats {
		stats.mu.RLock()
		for _, latency := range stats.Latencies {
			if latency > threshold {
				count++
			}
		}
		stats.mu.RUnlock()
	}
	return count
}

func (lm *LatencyManager) Reset() {
	lm.mu.Lock()
	defer lm.mu.Unlock()
	lm.stats = make(map[string]*LatencyStats)
}

// 全局延迟管理器实例
var GlobalLatencyManager = NewLatencyManager()
