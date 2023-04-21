import {
  Button,
  Flex,
  Icon,
  Input,
  Text,
  useColorModeValue,
} from "@chakra-ui/react";
import React, { useState } from "react";
import { useDropzone } from "react-dropzone";
import { AiOutlineFile } from "react-icons/ai";
import PropTypes from 'prop-types';

function Dropzone(props) {
  const { content, ...rest } = props;
  const [selectedFile, setSelectedFile] = useState(null);
  const { getRootProps, getInputProps } = useDropzone({
    onDrop: (acceptedFiles) => {
      setSelectedFile(acceptedFiles[0]);
    },
  });

  const bg = useColorModeValue("gray.100", "navy.700");
  const borderColor = useColorModeValue("secondaryGray.100", "whiteAlpha.100");
  const textColor = useColorModeValue("black", "white");

  return (
    <Flex
      align="center"
      justify="center"
      bg={bg}
      border="1px dashed"
      borderColor={borderColor}
      borderRadius="16px"
      w="100%"
      h="max-content"
      minH="100%"
      cursor="pointer"
      p={4}
      {...getRootProps({ className: "dropzone" })}
      {...rest}
    >
      <Input variant="main" {...getInputProps()} />
      {selectedFile ? (
        <Flex align="center" direction="column">
          <Icon
            as={AiOutlineFile}
            boxSize={12}
            color={textColor}
            mb={2}
          />
          <Text fontSize="sm" fontWeight="bold" color={textColor}>
            {selectedFile.name}
          </Text>
        </Flex>
      ) : (
        <Button
          variant="no-effects"
          borderColor={borderColor}
          color={textColor}
        >
          {content}
        </Button>
      )}
    </Flex>
  );
}

Dropzone.propTypes = {
  content: PropTypes.any
}

export default Dropzone;
