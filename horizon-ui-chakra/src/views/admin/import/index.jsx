// Chakra imports
import { Box, Grid } from "@chakra-ui/react";

// Custom components
import Upload from "views/admin/import/components/Upload";

// Assets
import React from "react";

export default function Overview() {
  return (
    <Box pt={{ base: "130px", md: "80px", xl: "80px" }}>
      <Grid>
        <Upload
          minH={{ base: "auto", lg: "420px", "2xl": "365px" }}
          pb={{ base: "100px", lg: "20px" }}
        />
      </Grid>
    </Box>
  );
}
