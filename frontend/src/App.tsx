import React from "react";
import { Board } from "@components/board";
import { Toolbar } from "@components/toolbar";
import { Box } from "@radix-ui/themes";

import { useData } from "./store";
import { Header } from "./views/layout/header";

export const App: React.FC = () => {
  const data = useData();

  return (
    <Box>
      <Header />
      <Toolbar />
      <Board {...data} />
    </Box>
  );
};
