/**
 * IP/CIDR validation utilities
 * Port of internal/utils/validation.go
 */

/**
 * Validates if a string is a valid IPv4 address
 */
export function isValidIPv4(ip: string): boolean {
  const parts = ip.split('.');
  if (parts.length !== 4) return false;
  
  for (const part of parts) {
    const num = parseInt(part, 10);
    if (isNaN(num) || num < 0 || num > 255) return false;
    // Check for leading zeros (invalid in strict IP notation)
    if (part !== String(num)) return false;
  }
  
  return true;
}

/**
 * Validates if a string is a valid IPv6 address
 */
export function isValidIPv6(ip: string): boolean {
  // Simple IPv6 validation - check for valid hex groups separated by colons
  const parts = ip.split(':');
  if (parts.length < 3 || parts.length > 8) return false;
  
  // Handle :: abbreviation
  const emptyParts = parts.filter(p => p === '').length;
  if (emptyParts > 1 && !(emptyParts === 2 && ip.includes('::'))) return false;
  
  for (const part of parts) {
    if (part === '') continue; // Allow empty parts for ::
    if (part.length > 4) return false;
    if (!/^[0-9a-fA-F]+$/.test(part)) return false;
  }
  
  return true;
}

/**
 * Validates if a string is a valid IP address (IPv4 or IPv6)
 */
export function isValidIP(ip: string): boolean {
  return isValidIPv4(ip) || isValidIPv6(ip);
}

/**
 * Validates if a string is a valid CIDR block
 */
export function isValidCIDR(cidr: string): boolean {
  const parts = cidr.split('/');
  if (parts.length !== 2) return false;
  
  const [ip, prefixStr] = parts;
  if (!ip || !prefixStr) return false;
  
  const prefix = parseInt(prefixStr, 10);
  if (isNaN(prefix)) return false;
  
  // IPv4 CIDR
  if (isValidIPv4(ip)) {
    return prefix >= 0 && prefix <= 32;
  }
  
  // IPv6 CIDR
  if (isValidIPv6(ip)) {
    return prefix >= 0 && prefix <= 128;
  }
  
  return false;
}

/**
 * Validates if a string is a valid IP address or CIDR block
 * Equivalent to Go's IsValidIPOrCIDR
 */
export function isValidIPOrCIDR(s: string): boolean {
  return isValidIP(s) || isValidCIDR(s);
}
