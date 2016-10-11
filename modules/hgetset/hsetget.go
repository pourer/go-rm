package main

import (
    "github.com/wenerme/go-rm/rm"
    "fmt"
    "github.com/urfave/cli"
)

func main() {
    fmt.Println("Enter main")
    app := cli.NewApp()
    _ = app
}

func init() {
    rm.Mod = CreateMyMod()
}
func CreateMyMod() *rm.Module {
    mod := rm.NewMod()
    mod.Name = "HelloWorld"
    mod.Version = 1

    mod.Commands = []rm.Command{
        {
            //
            // HGETSET key field value
            // Sets the 'field' in Hash 'key' to 'value' and returns the previous value, if
            // any.
            // Reply: String, the previous value or NULL if 'field' didn't exist.
            //
            Name:   "hsetget",
            Flags:  "write fast deny-oom",
            FirstKey:1, LastKey:1, KeyStep:1,
            Action: func(cmd rm.CmdContext) int {
                ctx := cmd.Ctx
                if len(cmd.Args) != 4 {
                    return ctx.WrongArity()
                }
                ctx.AutoMemory()

                // open the key and make sure it is indeed a Hash and not empty
                key := ctx.OpenKey(cmd.Args[1], rm.READ | rm.WRITE)
                if key.KeyType() != rm.KEYTYPE_EMPTY && key.KeyType() != rm.KEYTYPE_HASH {
                    ctx.ReplyWithError(rm.ERRORMSG_WRONGTYPE)
                    return rm.ERR
                }

                // get the current value of the hash element
                var val rm.String;
                key.HashGet(rm.HASH_NONE, cmd.Args[2], (*uintptr)(&val), rm.NullString())
                // set the element to the new value
                key.HashSet(rm.HASH_NONE, cmd.Args[2], cmd.Args[3], rm.NullString())

                if val.IsNull() {
                    ctx.ReplyWithNull()
                } else {
                    ctx.ReplyWithString(val)
                }
                return rm.OK
            },
        },
    }
    return mod
}