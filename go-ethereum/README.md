# EVM Extension (EthVM+)

This yields a new vm for Ethereum, handling new customized keywords. 
Current version of EthVM+ is made from go-ethereum. Thus for classical ethereum usage, please refer to https://github.com/ethereum/go-ethereum.


Following is the simplest example of the current version of EthVM+, with the new keyword 0x25 (SETNONFALLBACK);
<pre> <code>
escho@ubuntu:~/developments/go/src/github.com/ethereum/go-ethereum$ $GOPATH/bin/evm --debug --code 25  run 
0x
#### TRACE ####
SETNONFALLBACK  pc=00000000 gas=10000000000 cost=0

STOP            pc=00000001 gas=10000000000 cost=0

#### LOGS ####
</code> </pre>


The list of the modified Golang source files of go-ethereum are;
1. core/vm/opcodes.go
2. core/vm/instructions.go
3. core/vm/gas_table.go
4. core/vm/jump_table.go

If you have questions about technical details, please contact us.


As for Go-ethereum : Please see the [Developers' Guide](https://github.com/ethereum/go-ethereum/wiki/Developers'-Guide)
for more details on configuring your environment, managing project dependencies and testing procedures.

## License

The go-ethereum library (i.e. all code outside of the `cmd` directory) is licensed under the
[GNU Lesser General Public License v3.0](https://www.gnu.org/licenses/lgpl-3.0.en.html), also
included in our repository in the `COPYING.LESSER` file.

The go-ethereum binaries (i.e. all code inside of the `cmd` directory) is licensed under the
[GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html), also included
in our repository in the `COPYING` file.
