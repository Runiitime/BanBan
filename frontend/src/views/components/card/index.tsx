import React from "react";
import {
  Badge,
  Box,
  Card as CardUI,
  Flex,
  Heading,
  Text,
} from "@radix-ui/themes";

import type { ICard } from "@data/board";

import * as Styled from "./style.ts";

interface Props {
  data: ICard;
}

export const Card: React.FC<Props> = ({ data }: Props) => {
  return (
    <CardUI style={Styled.Card}>
      <Box>
        <Heading size="4">{data.title}</Heading>
      </Box>
      <Box>
        <Text>{data?.description}</Text>
      </Box>
      <Box>
        <Flex gap="2">
          <Badge color="ruby">Test</Badge>
          <Badge color="bronze">Tdse</Badge>
          <Badge color="amber">Vde</Badge>
        </Flex>
      </Box>
    </CardUI>
  );
};
