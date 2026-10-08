import type { CSSProperties } from "react";

export const Column: CSSProperties = {
  minWidth: "20%",
  border: "1px solid var(--gray-a4)",
  borderRadius: "var(--radius-4)",
  height: "calc(100vh - 124px",
  boxSizing: "border-box",
};

export const Heading: CSSProperties = {
  background: "var(--gray-a2)",
  borderBottom: "1px solid var(--gray-a4)",
};
