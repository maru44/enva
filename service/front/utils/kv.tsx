import { Kv } from '../types/kv'

// sort by key
export const sortKvs = (original: Kv[]): Kv[] => {
  return original.sort((kv1, kv2) => {
    if (kv1.kv_key > kv2.kv_key) {
      return 1
    }

    if (kv1.kv_key < kv2.kv_key) {
      return -1
    }

    return 0
  })
}

// @TODO sort by many pattern
