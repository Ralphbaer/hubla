// Chakra imports
import { Portal, Box, useDisclosure } from "@chakra-ui/react";
import Footer from "components/footer/FooterAdmin.js";
// Layout components
import Navbar from "components/navbar/NavbarAdmin.js";
import Sidebar from "components/sidebar/Sidebar.js";
import { SidebarContext } from "contexts/SidebarContext";
import React, { useState } from "react";
import { Redirect, Route, Switch } from "react-router-dom";
import routes from "routes.js";

// Custom Chakra theme
export default function Dashboard(props) {
  const { ...rest } = props;
  // states and functions
  const [fixed] = useState(false);
  const [toggleSidebar, setToggleSidebar] = useState(false);

  const isExcludedRoute = (route) => {
    return (
      (route.layout === "/auth" && route.path === "/sign-in")
      // (route.layout === "/admin" && route.path === "/profile")
    );
  }

  const filteredRoutes = routes.filter(route => !isExcludedRoute(route));

  // functions for changing the states from components
  const getRoute = () => {
    return window.location.pathname !== "/admin/full-screen-maps";
  };

  const getActiveRoute = (routes) => {
    const activeRoute = "Default Brand Text";
    for (const route of routes) {
      if (route.collapse) {
        const collapseActiveRoute = getActiveRoute(route.items);
        if (collapseActiveRoute !== activeRoute) {
          return collapseActiveRoute;
        }
      } else if (route.category) {
        const categoryActiveRoute = getActiveRoute(route.items);
        if (categoryActiveRoute !== activeRoute) {
          return categoryActiveRoute;
        }
      } else {
        if (
          window.location.href.indexOf(route.layout + route.path) !== -1
        ) {
          return route.name;
        }
      }
    }
    return activeRoute;
  };
  const getActiveNavbar = (routes) => {
    const activeNavbar = false;
    for (const route of routes) {
      if (route.collapse) {
        const collapseActiveNavbar = getActiveNavbar(route.items);
        if (collapseActiveNavbar !== activeNavbar) {
          return collapseActiveNavbar;
        }
      } else if (route.category) {
        const categoryActiveNavbar = getActiveNavbar(route.items);
        if (categoryActiveNavbar !== activeNavbar) {
          return categoryActiveNavbar;
        }
      } else if (
        window.location.href.includes(route.layout + route.path)
      ) {
        return route.secondary;
      }
    }
    return activeNavbar;
  };
  const getActiveNavbarText = (routes) => {
    const activeNavbar = false;
    for (const route of routes) {
      if (route.collapse) {
        const collapseActiveNavbar = getActiveNavbarText(route.items);
        if (collapseActiveNavbar !== activeNavbar) {
          return collapseActiveNavbar;
        }
      } else if (route.category) {
        const categoryActiveNavbar = getActiveNavbarText(route.items);
        if (categoryActiveNavbar !== activeNavbar) {
          return categoryActiveNavbar;
        }
      } else if (
        window.location.href.includes(route.layout + route.path)
      ) {
        return route.messageNavbar;
      }
    }
    return activeNavbar;
  };
  const getRoutes = (routes) => {
    return routes.map((prop, key) => {
      if (prop.layout === "/admin" && !isExcludedRoute(prop)) {
        return (
          <Route
            path={prop.layout + prop.path}
            component={prop.component}
            key={key}
          />
        );
      }
      if (prop.collapse) {
        return getRoutes(prop.items);
      }
      return prop.category ? getRoutes(prop.items) : null;
    });
  };
  document.documentElement.dir = "ltr";
  const { onOpen } = useDisclosure();
  return (
    <Box>
      <SidebarContext.Provider
        value={{
          toggleSidebar,
          setToggleSidebar,
        }}>
        <Sidebar routes={filteredRoutes} display='none' {...rest} />
        <Box
          float='right'
          minHeight='100vh'
          height='100%'
          overflow='auto'
          position='relative'
          maxHeight='100%'
          w={{ base: "100%", xl: "calc( 100% - 290px )" }}
          maxWidth={{ base: "100%", xl: "calc( 100% - 290px )" }}
          transition='all 0.33s cubic-bezier(0.685, 0.0473, 0.346, 1)'
          transitionDuration='.2s, .2s, .35s'
          transitionProperty='top, bottom, width'
          transitionTimingFunction='linear, linear, ease'>
          <Portal>
            <Box>
              <Navbar
                onOpen={onOpen}
                logoText={"Hubla Admin"}
                brandText={getActiveRoute(filteredRoutes)}
                secondary={getActiveNavbar(filteredRoutes)}
                message={getActiveNavbarText(filteredRoutes)}
                fixed={fixed}
                {...rest}
              />
            </Box>
          </Portal>

          {getRoute() ? (
            <Box
              mx='auto'
              p={{ base: "20px", md: "30px" }}
              pe='20px'
              minH='100vh'
              pt='50px'>
              <Switch>
                {getRoutes(filteredRoutes)}
                <Redirect from='/' to='/admin/default' />
              </Switch>
            </Box>
          ) : null}
          <Box>
            <Footer />
          </Box>
        </Box>
      </SidebarContext.Provider>
    </Box>
  );
}
