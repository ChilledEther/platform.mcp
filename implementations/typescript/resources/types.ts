
export interface ResourceContext {
  executionId: string;
}

export type ResourceContent = 
  | { uri: string; mimeType?: string; text: string; blob?: never }
  | { uri: string; mimeType?: string; blob: string; text?: never };

export interface ResourceResult {
  contents: ResourceContent[];
  isError?: boolean;
  meta?: Record<string, unknown>;
}

export type ResourceHandler = (
  uri: URL,
  context: ResourceContext
) => Promise<ResourceResult>;

export interface ResourceDefinition {
  name: string;
  uri: string;
  mimeType: string;
  description?: string;
  handler: ResourceHandler;
}
