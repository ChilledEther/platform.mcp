import { describe, expect, test } from 'bun:test';
import { isValidIPv4, isValidIPv6, isValidIP, isValidCIDR, isValidIPOrCIDR } from '../../utils/validation';

describe('isValidIPv4', () => {
  test('accepts valid IPv4 addresses', () => {
    expect(isValidIPv4('192.168.1.1')).toBe(true);
    expect(isValidIPv4('10.0.0.0')).toBe(true);
    expect(isValidIPv4('255.255.255.255')).toBe(true);
    expect(isValidIPv4('0.0.0.0')).toBe(true);
    expect(isValidIPv4('172.16.0.1')).toBe(true);
  });

  test('rejects invalid IPv4 addresses', () => {
    expect(isValidIPv4('256.1.1.1')).toBe(false);
    expect(isValidIPv4('192.168.1')).toBe(false);
    expect(isValidIPv4('192.168.1.1.1')).toBe(false);
    expect(isValidIPv4('192.168.1.a')).toBe(false);
    expect(isValidIPv4('')).toBe(false);
    expect(isValidIPv4('192.168.01.1')).toBe(false); // Leading zeros
  });
});

describe('isValidIPv6', () => {
  test('accepts valid IPv6 addresses', () => {
    expect(isValidIPv6('2001:0db8:85a3:0000:0000:8a2e:0370:7334')).toBe(true);
    expect(isValidIPv6('2001:db8:85a3::8a2e:370:7334')).toBe(true);
    expect(isValidIPv6('::1')).toBe(true);
    expect(isValidIPv6('fe80::1')).toBe(true);
  });

  test('rejects invalid IPv6 addresses', () => {
    expect(isValidIPv6('192.168.1.1')).toBe(false);
    expect(isValidIPv6('gggg::1')).toBe(false);
    expect(isValidIPv6('')).toBe(false);
  });
});

describe('isValidIP', () => {
  test('accepts both IPv4 and IPv6', () => {
    expect(isValidIP('192.168.1.1')).toBe(true);
    expect(isValidIP('::1')).toBe(true);
  });

  test('rejects invalid addresses', () => {
    expect(isValidIP('invalid')).toBe(false);
    expect(isValidIP('')).toBe(false);
  });
});

describe('isValidCIDR', () => {
  test('accepts valid IPv4 CIDR blocks', () => {
    expect(isValidCIDR('10.0.0.0/8')).toBe(true);
    expect(isValidCIDR('192.168.1.0/24')).toBe(true);
    expect(isValidCIDR('0.0.0.0/0')).toBe(true);
    expect(isValidCIDR('192.168.1.1/32')).toBe(true);
  });

  test('accepts valid IPv6 CIDR blocks', () => {
    expect(isValidCIDR('2001:db8::/32')).toBe(true);
    expect(isValidCIDR('::1/128')).toBe(true);
  });

  test('rejects invalid CIDR blocks', () => {
    expect(isValidCIDR('192.168.1.1')).toBe(false); // No prefix
    expect(isValidCIDR('192.168.1.1/33')).toBe(false); // Invalid IPv4 prefix
    expect(isValidCIDR('::1/129')).toBe(false); // Invalid IPv6 prefix
    expect(isValidCIDR('192.168.1.1/')).toBe(false);
    expect(isValidCIDR('/24')).toBe(false);
    expect(isValidCIDR('')).toBe(false);
  });
});

describe('isValidIPOrCIDR', () => {
  test('accepts IPs and CIDRs', () => {
    expect(isValidIPOrCIDR('192.168.1.1')).toBe(true);
    expect(isValidIPOrCIDR('10.0.0.0/8')).toBe(true);
    expect(isValidIPOrCIDR('::1')).toBe(true);
    expect(isValidIPOrCIDR('2001:db8::/32')).toBe(true);
  });

  test('rejects invalid inputs', () => {
    expect(isValidIPOrCIDR('invalid')).toBe(false);
    expect(isValidIPOrCIDR('')).toBe(false);
    expect(isValidIPOrCIDR('192.168.1.1/33')).toBe(false);
  });
});
