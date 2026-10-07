import React from "react";
import { Box, Flex, Heading } from "@radix-ui/themes";
import { map } from "lodash";

import type { ICard, IColumn } from "@data/board";

import { Card } from "../card";
import * as Styled from "./style";

interface Props {
  data: IColumn;
}

export const Column: React.FC<Props> = ({ data }: Props) => {
  const { cards, title } = data;

  const cardsItems = React.useMemo(() => {
    return map(cards, (item: ICard) => <Card data={item} />);
  }, [cards]);

  return (
    <Box style={Styled.Column}>
      <Box py="3" mb="4" style={Styled.Heading}>
        <Heading>{title}</Heading>
      </Box>
      <Flex direction={"column"} gap={"3"}>
        {cardsItems}
      </Flex>
    </Box>
  );
};
