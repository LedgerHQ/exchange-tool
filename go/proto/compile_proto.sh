# SPDX-FileCopyrightText: 2024 Ledger
#
# SPDX-License-Identifier: Apache-2.0

#!/bin/sh

#protoc --proto_path=./proto --go_out=./proto --go_out=Mproto/protocol.proto=swap.ledger.fr/proto ./proto/protocol.proto
protoc --go_out=. ./protocol.proto -I .