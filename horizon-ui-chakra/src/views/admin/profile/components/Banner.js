// Chakra imports
import { Avatar, Box, Flex, Text, useColorModeValue } from "@chakra-ui/react";
import Card from "components/card/Card.js";
import React from "react";

export default function Banner(props) {
  const { banner, avatar, name, creatorType, balance, createdAt, updatedAt } = props;
  // Chakra Color Mode
  const textColorPrimary = useColorModeValue("secondaryGray.900", "white");
  const textColorSecondary = "gray.400";
  const borderColor = useColorModeValue(
    "white !important",
    "#111C44 !important"
  );
  return (
    <Card mb={{ base: "0px", lg: "20px" }} align='center'>
      <Box
        bg={`url(${banner})`}
        bgSize='cover'
        borderRadius='16px'
        h='250px'
        w='100%'
      />
      <Avatar
        mx='auto'
        src={avatar}
        h='87px'
        w='87px'
        mt='-43px'
        border='4px solid'
        borderColor={borderColor}
      />
      <Text color={textColorPrimary} fontWeight='bold' fontSize='xl' mt='10px'>
        {name}
      </Text>
      <Text color={textColorSecondary} fontSize='sm'>
        {creatorType}
      </Text>
      <Flex w='100%' justify='space-between' mx='auto' mt='26px'>
        <Flex ml='80px' align='center' direction='column'>
          <Text color={textColorSecondary} fontSize='sm' fontWeight='400'>
            Data de Criação
          </Text>
          <Text color={textColorPrimary} fontSize='sm' fontWeight='700'>
            {createdAt}
          </Text>
        </Flex>
        <Flex pl='20px' align='center' direction='column' flex='1' justify='center'>
          <Text color={textColorSecondary} fontSize='sm' fontWeight='400'>
            Total em Vendas
          </Text>
          <Text color={textColorPrimary} fontSize='2xl' fontWeight='700'>
            R${balance}
          </Text>
        </Flex>
        <Flex mr='80px' align='center' direction='column'>
          <Text color={textColorSecondary} fontSize='sm' fontWeight='400'>
            Última Atualização
          </Text>
          <Text color={textColorPrimary} fontSize='sm' fontWeight='700'>
            {updatedAt}
          </Text>
        </Flex>
      </Flex>
    </Card >
  );
}
