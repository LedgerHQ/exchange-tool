/*
 *  Licensed to the Apache Software Foundation (ASF) under one
 *  or more contributor license agreements.  See the NOTICE file
 *  distributed with this work for additional information
 *  regarding copyright ownership.  The ASF licenses this file
 *  to you under the Apache License, Version 2.0 (the
 *  "License"); you may not use this file except in compliance
 *  with the License.  You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing,
 *  software distributed under the License is distributed on an
 *  "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 *  KIND, either express or implied.  See the License for the
 *  specific language governing permissions and limitations
 *  under the License.
 */

package main

import (
	"swap.ledger.fr/cmd"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "exchange-tool",
	Short: "A tool for generated test data for exchange features and test it",
	Long: `ExchangeTool is CLI for helping Exchange Developer and Ledger's partner.
It can generate test data to use in LiveApp to test app-exchange signature.
It can check test data generate by provider to verify format and signature`,
}

func main() {
	rootCmd.AddCommand(
		cmd.GenerateCmd,
		cmd.CheckCmd,
		cmd.ReadCmd,
		cmd.HexCmd,
		cmd.SignCmd,
		cmd.CalCmd,
	)

	rootCmd.Execute()
}
