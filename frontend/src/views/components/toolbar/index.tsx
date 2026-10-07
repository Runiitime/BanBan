import React from "react";
import { Box, Button, Flex } from "@radix-ui/themes";

import * as Styled from "./style";

export const Toolbar: React.FC = () => {
  return (
    <Box p="3" style={Styled.Toolbar}>
      <Flex gap="4">
        <Button>Add</Button>
        <Button>Find</Button>
        <Button>Delete</Button>
      </Flex>
    </Box>
  );
};
