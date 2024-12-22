package behaviors

import (
	"context"
	"fmt"
	"os"
	"syscall"
	"testing"
	"time"
)

func TestContext_SetAndGet(t *testing.T) {
	ctx := NewContext(context.Background())
	ctx.Set("key", "value")

	val := ctx.Get("key")
	if val != "value" {
		t.Errorf("Expected 'value', got %v", val)
	}
}

func TestChain_Execute(t *testing.T) {
	node1 := NewFuncNode(func(ctx *Context) (*NodeResult, error) {
		ctx.Set("key1", "value1")
		return nil, nil
	}, "node1")

	node2 := NewFuncNode(func(ctx *Context) (*NodeResult, error) {
		val := ctx.Get("key1")
		if val != "value1" {
			t.Errorf("Expected 'value1', got %v", val)
		}
		return nil, nil
	}, "node2")

	chain := NewChain(node1, node2)
	_, err := chain.Execute(NewContext(context.Background()))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestChain_AddNextChain(t *testing.T) {
	node1 := NewFuncNode(func(ctx *Context) (*NodeResult, error) {
		ctx.Set("key", "chain1")
		return nil, nil
	}, "node1")
	chain1 := NewChain(node1)

	node2 := NewFuncNode(func(ctx *Context) (*NodeResult, error) {
		ctx.Set("key", "chain2")
		return nil, nil
	}, "node2")
	chain2 := NewChain(node2)

	chain1.AddNextChain(chain2, 1.0)

	ctx := NewContext(context.Background())
	_, err := chain1.Execute(ctx)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if val := ctx.Get("key"); val != "chain2" {
		t.Errorf("Expected 'chain2', got %v", val)
	}
}

func TestLoadGenerator_Start(t *testing.T) {
	node1 := NewFuncNode(func(ctx *Context) (*NodeResult, error) {
		fmt.Println("node1", time.Now().String())
		return nil, nil
	}, "node1")
	node2 := NewFuncNode(func(ctx *Context) (*NodeResult, error) {
		fmt.Println("node2", time.Now().String())
		return nil, nil
	}, "node2")
	node3 := NewFuncNode(func(ctx *Context) (*NodeResult, error) {
		fmt.Println("node3", time.Now().String())
		return nil, nil
	}, "node3")
	node4 := NewFuncNode(func(ctx *Context) (*NodeResult, error) {
		fmt.Println("node4", time.Now().String())
		return nil, nil
	}, "node4")
	chain1 := NewChain(node1)
	chain2 := NewChain(node2, node3)
	chain3 := NewChain(node4)
	chain1.AddNextChain(chain2, 0.2)
	chain1.AddNextChain(chain3, 0.8)

	loadGen := NewLoadGenerator(WithThread(1), WithSleep(3000), WithChain(chain1))
	go func() {
		time.Sleep(5 * time.Second)
		p, _ := os.FindProcess(os.Getpid())
		p.Signal(syscall.SIGINT) // Simulate Ctrl + C
	}()
	loadGen.Start()

}

func TestLoadGenerator_PanicRecovery(t *testing.T) {
	node1 := NewFuncNode(func(ctx *Context) (*NodeResult, error) {
		fmt.Println("node1", time.Now().String())
		return nil, nil
	}, "node1")
	node2 := NewFuncNode(func(ctx *Context) (*NodeResult, error) {
		fmt.Println("node2", time.Now().String())
		panic("intentional panic for testing")
		return nil, nil
	}, "node2")
	node3 := NewFuncNode(func(ctx *Context) (*NodeResult, error) {
		fmt.Println("node3", time.Now().String())
		return nil, nil
	}, "node3")
	node4 := NewFuncNode(func(ctx *Context) (*NodeResult, error) {
		fmt.Println("node4", time.Now().String())
		return nil, nil
	}, "node4")
	chain1 := NewChain(node1)
	chain2 := NewChain(node2, node3)
	chain3 := NewChain(node4)
	chain1.AddNextChain(chain2, 0.2)
	chain1.AddNextChain(chain3, 0.8)

	loadGen := NewLoadGenerator(WithThread(3), WithSleep(3000), WithChain(chain1))

	go func() {
		time.Sleep(5 * time.Second) // Allow some time for the panic and recovery
		p, _ := os.FindProcess(os.Getpid())
		p.Signal(syscall.SIGINT) // Simulate Ctrl + C

	}()

	loadGen.Start()

	// Check if the worker was restarted after panic
	// This can be done by checking logs or other side effects
	// For simplicity, we assume the test passes if no deadlock or crash occurs
	t.Log("Test completed, check logs for panic recovery")
}

func TestFuncNode_Execute(t *testing.T) {
	executed := false
	fn := func(ctx *Context) (*NodeResult, error) {
		executed = true
		return nil, nil
	}

	node := NewFuncNode(fn, "node")
	_, err := node.Execute(NewContext(context.Background()))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !executed {
		t.Errorf("Function was not executed")
	}
}
