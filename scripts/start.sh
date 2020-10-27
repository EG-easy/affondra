#!/usr/bin/env bash

rm -rf ~/.affondrad
rm -rf ~/.affondracli

affondrad init test --chain-id=affondra

affondracli config output json
affondracli config indent true
affondracli config trust-node true
affondracli config chain-id affondra
affondracli config keyring-backend test

affondracli keys add user1
affondracli keys add user2
affondracli keys add faucet

affondrad add-genesis-account $(affondracli keys show user1 -a) 1000token,100000000stake
affondrad add-genesis-account $(affondracli keys show user2 -a) 1000token,100000000stake
affondrad add-genesis-account $(affondracli keys show faucet -a) 100000000stake

affondrad gentx --name user1 --keyring-backend test

echo "Collecting genesis txs..."
affondrad collect-gentxs

echo "Validating genesis file..."
affondrad validate-genesis
