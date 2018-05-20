package main

import (
	"fmt"

	"github.com/rodrigodiez/zorro/pkg/generator/uuid"
	"github.com/rodrigodiez/zorro/pkg/service"
	"github.com/rodrigodiez/zorro/pkg/storage/memory"
)

func main() {
	z := service.New(
		uuid.NewV4(),
		memory.New(),
	)

	value := z.Mask("foo")
	fmt.Println(value)
	// Will print something like '870284f9-c269-4175-8ab9-8e0a094a64ab'

	key, _ := z.Unmask(value)
	fmt.Println(key)
	// Will print 'foo'

	// Once generated masks are idempotent!
	value = z.Mask("foo")
	fmt.Println(value)
	// Will print same mask as before
}
