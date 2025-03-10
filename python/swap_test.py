# SPDX-FileCopyrightText: Ledger SAS 2024
#
# SPDX-License-Identifier: Apache-2.0

from cryptography.hazmat.primitives.asymmetric import ec
from swap import sign_payload, check_signature

test_data = [{
    "name": "Auto Ledger",
    "payload": "CipiYzFxYXIwc3Jycjd4Zmt2eTVsNjQzbHlkbnc5cmU1OWd0enp3ZjVtZHEaKmJjMXFhcjBzcnJyN3hma3Z5NWw2NDNseWRudzlyZTU5Z3R6endmNHRlcSoqMHhiNzk0ZjVlYTBiYTM5NDk0Y2U4Mzk2MTNmZmZiYTc0Mjc5NTc5MjY4OgNCVENCA0JBVEoDr3ngUggIU6DSMTwAAFoKQUJDREVGR0hJSmIKQUJDREVGR0hJSg==",
    "signature": None,
    "priv": "../samples/sample-priv-key-secp256k1.pem",
    "pub": "../samples/sample-pub-key-secp256k1.pem",
    "curve": None
  }, {
    "name": "Auto Check R1",
    "payload": "CipiYzFxYXIwc3Jycjd4Zmt2eTVsNjQzbHlkbnc5cmU1OWd0enp3ZjVtZHEaKmJjMXFhcjBzcnJyN3hma3Z5NWw2NDNseWRudzlyZTU5Z3R6endmNHRlcSoqMHhiNzk0ZjVlYTBiYTM5NDk0Y2U4Mzk2MTNmZmZiYTc0Mjc5NTc5MjY4OgNCVENCA0JBVEoDr3ngUggIU6DSMTwAAFoKQUJDREVGR0hJSmIKQUJDREVGR0hJSg==",
    "signature": None,
    "priv": "../samples/sample-priv-key-secp256r1.pem",
    "pub": "../samples/sample-pub-key-secp256r1.pem",
    "curve": ec.SECP256R1
  }, {
    "name": "Check R1",
    "payload": "CipiYzFxYXIwc3Jycjd4Zmt2eTVsNjQzbHlkbnc5cmU1OWd0enp3ZjVtZHEaKmJjMXFhcjBzcnJyN3hma3Z5NWw2NDNseWRudzlyZTU5Z3R6endmNHRlcSoqMHhiNzk0ZjVlYTBiYTM5NDk0Y2U4Mzk2MTNmZmZiYTc0Mjc5NTc5MjY4OgNCVENCA0JBVEoDr3ngUggIU6DSMTwAAFoKQUJDREVGR0hJSmIKQUJDREVGR0hJSg==",
    "signature": "rrSlSeyUo3XRBRAKqqTjsxXb-_oif4_JnXaUzDVMyhI0M_5AOdmYbDfdVOTb41nIrxQGxdnapULRtPQoXeKbnA==", # Signature comes from Go code
    "pub": "../samples/sample-pub-key-secp256r1.pem",
    "curve": ec.SECP256R1
  }, {
    "name": "Auto Check K1",
    "payload": ".CipiYzFxYXIwc3Jycjd4Zmt2eTVsNjQzbHlkbnc5cmU1OWd0enp3ZjVtZHEaKmJjMXFhcjBzcnJyN3hma3Z5NWw2NDNseWRudzlyZTU5Z3R6endmNHRlcSoqMHhiNzk0ZjVlYTBiYTM5NDk0Y2U4Mzk2MTNmZmZiYTc0Mjc5NTc5MjY4OgNCVENCA0JBVEoDr3ngUggIU6DSMTwAAFoKQUJDREVGR0hJSmIKQUJDREVGR0hJSg==",
    "signature": None,
    "priv": "../samples/sample-priv-key-secp256k1.pem",
    "pub": "../samples/sample-pub-key-secp256k1.pem",
    "curve": None
  }, {
    "name": "Check K1",
    "payload": "CipiYzFxYXIwc3Jycjd4Zmt2eTVsNjQzbHlkbnc5cmU1OWd0enp3ZjVtZHEaKmJjMXFhcjBzcnJyN3hma3Z5NWw2NDNseWRudzlyZTU5Z3R6endmNHRlcSoqMHhiNzk0ZjVlYTBiYTM5NDk0Y2U4Mzk2MTNmZmZiYTc0Mjc5NTc5MjY4OgNCVENCA0JBVEoDr3ngUggIU6DSMTwAAFoKQUJDREVGR0hJSmIKQUJDREVGR0hJSg==",
    "signature": "JkzwubY_waKOBzhF5tICHmjdU8egRpRZ4EnUPEwzK_MKzsdVYeJwJ2lCQqg2ZBA_QdjgDmMQ_-wNisHbVCe4dA==", # Signature comes from Go code
    "pub": "../samples/sample-pub-key-secp256k1.pem",
    "curve": ec.SECP256K1
}]

PRINT_COLOR_GREEN = '\033[92m'
PRINT_COLOR_RED = '\033[91m'
PRINT_COLOR_END = '\033[0m'

for data in test_data:
  name = data["name"]
  payload = data["payload"]
  signature_base64 = data["signature"]
  pub_key = data["pub"]
  curve = data["curve"]
  form = data.get("format", "raw")

  print(f"## {name} ##")

  if curve is None:
    curve = ec.SECP256K1
  if signature_base64 is None:
    priv_key = data["priv"]
    signature_base64 = sign_payload(priv_key, payload, curve)
    # print("Base64 signature:", signature_base64)

  try:
    check_signature(pub_key, payload, signature_base64, curve, form)
    print(f"{name} with its value {PRINT_COLOR_GREEN}SUCCEEDED{PRINT_COLOR_END}")
  except Exception as error:
    print(f"{name} with its value {PRINT_COLOR_RED}FAILED{PRINT_COLOR_END}", type(error).__name__, " - ", error)

  print("\n")
