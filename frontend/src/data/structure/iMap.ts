import type { UniqueIdentifier } from "@dnd-kit/abstract";

export type IMap<T> = Record<string | number, T>;

export type MapI<T> = Record<UniqueIdentifier, T>;
