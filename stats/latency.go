package stats

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

type LatencyRecord struct {
	Latency   time.Duration
	Timestamp time.Time
}

type LatencyStats struct {
	URL        string
	Method     string
	Records    []LatencyRecord
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
		Records:    make([]LatencyRecord, 0),
		MinLatency: time.Duration(^uint64(0) >> 1), // 最大值
	}
}

func (ls *LatencyStats) AddLatency(latency time.Duration) {
	ls.mu.Lock()
	defer ls.mu.Unlock()

	record := LatencyRecord{
		Latency:   latency,
		Timestamp: time.Now(),
	}
	ls.Records = append(ls.Records, record)
	ls.Count++

	if latency > ls.MaxLatency {
		ls.MaxLatency = latency
	}
	if latency < ls.MinLatency {
		ls.MinLatency = latency
	}

	// 计算平均延迟
	total := time.Duration(0)
	for _, record := range ls.Records {
		total += record.Latency
	}
	ls.AvgLatency = total / time.Duration(ls.Count)
}

func (ls *LatencyStats) GetPercentile(percentile float64) time.Duration {
	ls.mu.RLock()
	defer ls.mu.RUnlock()

	if len(ls.Records) == 0 {
		return 0
	}

	// 创建延迟切片并排序
	sortedLatencies := make([]time.Duration, len(ls.Records))
	for i, record := range ls.Records {
		sortedLatencies[i] = record.Latency
	}
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

	now := time.Now()
	timeWindow := time.Minute // 只检查最近1分钟

	for _, stats := range lm.stats {
		stats.mu.RLock()

		// 检查最近1分钟内的请求
		for i := len(stats.Records) - 1; i >= 0; i-- {
			record := stats.Records[i]

			// 如果请求时间超过1分钟，停止检查
			if now.Sub(record.Timestamp) > timeWindow {
				break
			}

			// 检查是否有慢请求
			if record.Latency > threshold {
				stats.mu.RUnlock()
				return true
			}
		}
		stats.mu.RUnlock()
	}
	return false
}

func (lm *LatencyManager) GetSlowRequestsCount(threshold time.Duration) int {
	lm.mu.RLock()
	defer lm.mu.RUnlock()

	count := 0
	for _, stats := range lm.stats {
		stats.mu.RLock()
		for _, record := range stats.Records {
			if record.Latency > threshold {
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

func (lm *LatencyManager) HasRecentSlowRequests(threshold time.Duration, timeWindow time.Duration) bool {
	lm.mu.RLock()
	defer lm.mu.RUnlock()

	now := time.Now()

	for _, stats := range lm.stats {
		stats.mu.RLock()

		for i := len(stats.Records) - 1; i >= 0; i-- {
			record := stats.Records[i]

			if now.Sub(record.Timestamp) > timeWindow {
				break
			}

			if record.Latency > threshold {
				stats.mu.RUnlock()
				return true
			}
		}
		stats.mu.RUnlock()
	}
	return false
}

// CleanOldRecords 清理超过指定时间的旧记录
func (ls *LatencyStats) CleanOldRecords(maxAge time.Duration) {
	ls.mu.Lock()
	defer ls.mu.Unlock()

	now := time.Now()
	cutoff := now.Add(-maxAge)

	// 找到第一个不需要删除的记录
	firstValidIndex := 0
	for i, record := range ls.Records {
		if record.Timestamp.After(cutoff) {
			firstValidIndex = i
			break
		}
	}

	// 如果所有记录都太旧，清空
	if firstValidIndex == 0 && len(ls.Records) > 0 && ls.Records[0].Timestamp.Before(cutoff) {
		ls.Records = ls.Records[:0]
		ls.Count = 0
		ls.MaxLatency = 0
		ls.MinLatency = time.Duration(^uint64(0) >> 1)
		ls.AvgLatency = 0
		return
	}

	// 删除旧记录
	if firstValidIndex > 0 {
		ls.Records = ls.Records[firstValidIndex:]
		ls.Count = len(ls.Records)

		// 重新计算统计数据
		ls.recalculateStats()
	}
}

// recalculateStats 重新计算统计数据
func (ls *LatencyStats) recalculateStats() {
	if len(ls.Records) == 0 {
		ls.MaxLatency = 0
		ls.MinLatency = time.Duration(^uint64(0) >> 1)
		ls.AvgLatency = 0
		return
	}

	ls.MaxLatency = 0
	ls.MinLatency = time.Duration(^uint64(0) >> 1)
	total := time.Duration(0)

	for _, record := range ls.Records {
		if record.Latency > ls.MaxLatency {
			ls.MaxLatency = record.Latency
		}
		if record.Latency < ls.MinLatency {
			ls.MinLatency = record.Latency
		}
		total += record.Latency
	}

	ls.AvgLatency = total / time.Duration(len(ls.Records))
}

// CleanOldRecords 清理所有统计中的旧记录
func (lm *LatencyManager) CleanOldRecords(maxAge time.Duration) {
	lm.mu.RLock()
	defer lm.mu.RUnlock()

	for _, stats := range lm.stats {
		stats.CleanOldRecords(maxAge)
	}
}
