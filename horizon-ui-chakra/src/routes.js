import React from "react";

import { Icon } from "@chakra-ui/react";
import {
  MdPerson,
  MdHome,
  MdLock,
  MdFilePresent,
  MdImportExport
} from "react-icons/md";

// Admin Imports
import MainDashboard from "views/admin/default";
import Profile from "views/admin/profile";
import Import from "views/admin/import";
import FileTransaction from "views/admin/file-transaction";

// Auth Imports
import SignInCentered from "views/auth/signIn";

const routes = [
  {
    name: "Menu Principal",
    layout: "/admin",
    path: "/default",
    icon: <Icon as={MdHome} width='20px' height='20px' color='inherit' />,
    component: MainDashboard,
  },
  {
    name: "Profile",
    layout: "/admin",
    path: "/profile",
    icon: <Icon as={MdPerson} width='20px' height='20px' color='inherit' />,
    component: Profile,
  },
  {
    name: "Sign In",
    layout: "/auth",
    path: "/sign-in",
    icon: <Icon as={MdLock} width='20px' height='20px' color='inherit' />,
    component: SignInCentered,
  },
  {
    name: "Importar Arquivo",
    layout: "/admin",
    path: "/import",
    icon: <Icon as={MdFilePresent} width='20px' height='20px' color='inherit' />,
    component: Import,
  },
  {
    name: "Transações Importadas",
    layout: "/admin",
    path: "/file-transaction",
    icon: <Icon as={MdImportExport} width='20px' height='20px' color='inherit' />,
    component: FileTransaction,
  },
];

export default routes;
