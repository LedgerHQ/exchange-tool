# SPDX-FileCopyrightText: Ledger SAS 2024
#
# SPDX-License-Identifier: Apache-2.0

import unittest
import base64
from cryptography.hazmat.primitives.asymmetric import ec
from swap import sign_payload, check_signature

class TestSwap(unittest.TestCase):
    def setUp(self):
        self.payload = "CipiYzFxYXIwc3Jycjd4Zmt2eTVsNjQzbHlkbnc5cmU1OWd0enp3ZjVtZHEaKmJjMXFhcjBzcnJyN3hma3Z5NWw2NDNseWRudzlyZTU5Z3R6endmNHRlcSoqMHhiNzk0ZjVlYTBiYTM5NDk0Y2U4Mzk2MTNmZmZiYTc0Mjc5NTc5MjY4OgNCVENCA0JBVEoDr3ngUggIU6DSMTwAAFoKQUJDREVGR0hJSmIKQUJDREVGR0hJSg=="
        self.k1_priv = "../samples/sample-priv-key-secp256k1.pem"
        self.k1_pub = "../samples/sample-pub-key-secp256k1.pem"
        self.r1_priv = "../samples/sample-priv-key-secp256r1.pem"
        self.r1_pub = "../samples/sample-pub-key-secp256r1.pem"

    def test_auto_ledger_k1(self):
        signature = sign_payload(self.k1_priv, self.payload, ec.SECP256K1)
        check_signature(self.k1_pub, self.payload, signature, ec.SECP256K1)

    def test_auto_check_r1(self):
        signature = sign_payload(self.r1_priv, self.payload, ec.SECP256R1)
        check_signature(self.r1_pub, self.payload, signature, ec.SECP256R1)

    def test_check_r1_fixed_signature(self):
        # Signature comes from Go code
        signature = "rrSlSeyUo3XRBRAKqqTjsxXb-_oif4_JnXaUzDVMyhI0M_5AOdmYbDfdVOTb41nIrxQGxdnapULRtPQoXeKbnA=="
        check_signature(self.r1_pub, self.payload, signature, ec.SECP256R1)

    def test_auto_check_k1_jwt_style(self):
        # The payload in the original test had a leading dot for K1 check
        payload_with_dot = "." + self.payload
        signature = sign_payload(self.k1_priv, payload_with_dot, ec.SECP256K1)
        check_signature(self.k1_pub, payload_with_dot, signature, ec.SECP256K1)

    def test_check_k1_fixed_signature(self):
        # Signature comes from Go code
        signature = "JkzwubY_waKOBzhF5tICHmjdU8egRpRZ4EnUPEwzK_MKzsdVYeJwJ2lCQqg2ZBA_QdjgDmMQ_-wNisHbVCe4dA=="
        check_signature(self.k1_pub, self.payload, signature, ec.SECP256K1)

if __name__ == '__main__':
    unittest.main()
