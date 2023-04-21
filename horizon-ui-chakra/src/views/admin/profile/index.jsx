// Chakra imports
import { Box, Grid } from "@chakra-ui/react";

// Custom components
import Banner from "views/admin/profile/components/Banner";

// Assets
import banner from "assets/img/auth/banner.png";
import React from "react";

export default function Overview() {
  return (
    <Box pt={{ base: "130px", md: "80px", xl: "80px" }}>
      <Grid
        gap={{ base: "20px", xl: "20px" }}>
        <Banner
          banner={banner}
          avatar='1' //make it dynamic
          name='Adela Parkson'
          creatorType='CREATOR'
          balance='70000'
          createdAt='03/03/2024'
          updatedAt='05/05/2024 às 12:00'
        />
      </Grid>
    </Box>
  );
}
