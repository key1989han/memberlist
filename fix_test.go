package memberlist

import "testing"

func TestFixedMemberlist_ConcurrentAccess(t *testing.T) {
    ml := &FixedMemberlist{
        nodeMap: make(map[string]*Node),
    }

    // Test concurrent writes
    done := make(chan bool)
    for i := 0; i < 10; i++ {
        go func(i int) {
            ml.SetNode("node"+string(rune('A'+i)), &Node{State: StateAlive})
            done <- true
        }(i)
    }
    for i := 0; i < 10; i++ {
        <-done
    }

    // Test concurrent reads
    for i := 0; i < 10; i++ {
        go func(i int) {
            ml.GetNode("node" + string(rune('A'+i)))
            done <- true
        }(i)
    }
    for i := 0; i < 10; i++ {
        <-done
    }
}

func TestFixedMemberlist_RemoveNode(t *testing.T) {
    ml := &FixedMemberlist{
        nodeMap: map[string]*Node{"test": {State: StateAlive}},
    }
    ml.RemoveNode("test")
    if ml.GetNode("test") != nil {
        t.Error("Expected nil after remove")
    }
}
