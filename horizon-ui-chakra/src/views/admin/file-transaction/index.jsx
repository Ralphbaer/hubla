// Chakra imports
import {
  Box,
  SimpleGrid,
} from "@chakra-ui/react";
// Custom components
import React from "react";
import IDTable from "views/admin/file-transaction/components/IDTable";
import {
  columnsDataID,
} from "views/admin/file-transaction/variables/columnsData";
import tableDataID from "views/admin/file-transaction/variables/tableDataID.json";

export default function UserReports() {
  // Chakra Color Mode
  return (
    <Box pt={{ base: "130px", md: "80px", xl: "80px" }}>
      <SimpleGrid gap='20px' mb='20px'>
        <IDTable columnsData={columnsDataID} tableData={tableDataID} />
      </SimpleGrid>
    </Box>
  );
}
