import PropTypes from "prop-types"
import { Box, useStyleConfig } from "@chakra-ui/react";
import React from 'react';

function Card(props) {
  const { variant, children, ...rest } = props;
  const styles = useStyleConfig("Card", { variant });

  return (
    <Box __css={styles} {...rest}>
      {children}
    </Box>
  );
}

Card.propTypes = {
  children: PropTypes.any,
  variant: PropTypes.any
}

export default Card;
