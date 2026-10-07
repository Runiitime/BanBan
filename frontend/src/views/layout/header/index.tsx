import React from "react";
import { Avatar, Box, Flex, Heading } from "@radix-ui/themes";

import * as Styled from "./style";

export const Header: React.FC = () => {
  return (
    <Box p="1" style={Styled.Header}>
      <Flex justify="between" align="center">
        <Heading size={"5"}>BanBan kanban</Heading>
        <Avatar
          size="2"
          src="https://images.unsplash.com/photo-1502823403499-6ccfcf4fb453?&w=256&h=256&q=70&crop=focalpoint&fp-x=0.5&fp-y=0.3&fp-z=1&fit=crop"
          fallback="A"
        />
      </Flex>
    </Box>
  );
};
