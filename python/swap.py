# SPDX-FileCopyrightText: 2024 Ledger
#
# SPDX-License-Identifier: Apache-2.0

import sys
import base64
from cryptography.hazmat.primitives import serialization, hashes
from cryptography.hazmat.primitives.asymmetric import ec
from cryptography.hazmat.primitives.asymmetric.utils import decode_dss_signature, encode_dss_signature

# Encode
def sign_payload(privkey_filename, payload, curve = ec.SECP256K1):
  ## Open Private key
  with open(privkey_filename, "rb") as f:
    private_key = serialization.load_pem_private_key(f.read(), password=None)

  assert isinstance(private_key.curve, curve)

  ## Sign Payload
  der_signature = private_key.sign(str.encode(payload), ec.ECDSA(hashes.SHA256()))
  r, s = decode_dss_signature(der_signature)
  r_s_signature = r.to_bytes(32, "big") + s.to_bytes(32, "big")

  signature_base64 = base64.urlsafe_b64encode(r_s_signature)

  return signature_base64

# Decode
def check_signature(pubkey_filename, payload, signature, curve = ec.SECP256K1, format = "raw"):
  ## Open Public key
  with open(pubkey_filename, "rb") as f:
    public_key = serialization.load_pem_public_key(f.read(), )

  assert isinstance(public_key.curve, curve)

  ## Check signature
  # public_key.verify(encode_dss_signature(r, s), str.encode(payload), ec.ECDSA(hashes.SHA256())) # --> OK
  decoded_signature = base64.urlsafe_b64decode(signature)
  r_s_decoded = (int.from_bytes(decoded_signature[:32], "big"), int.from_bytes(decoded_signature[32:], "big"))
  # print("R found:", decoded_signature[:32].hex())
  # print("S found:", decoded_signature[32:].hex())
  formated_payload = str.encode(payload)
  if format == "jwt":
    formated_payload = str.encode('.'+payload)

  public_key.verify(encode_dss_signature(r_s_decoded[0], r_s_decoded[1]), formated_payload, ec.ECDSA(hashes.SHA256()))
  # print("Signature used:", decoded_signature.hex())
  # print("Public key used:", public_key.public_numbers())
  # public_bytes = public_key.public_bytes(
  #   encoding=serialization.Encoding.X962,
  #   format=serialization.PublicFormat.UncompressedPoint,
  # )
  # print("Public key used (CAL format):", public_bytes.hex())

def main(argv) -> int:
  print("Main", argv[1])
  return 0

if __name__ == '__main__':
  sys.exit(main(sys.argv))
