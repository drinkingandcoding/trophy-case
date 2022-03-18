import React, { useRef, useEffect } from "react";
import { useLocation, Switch } from "react-router-dom";
import AppRoute from "./utils/AppRoute";
import ScrollReveal from "./utils/ScrollReveal";
import { QueryClient, QueryClientProvider } from "react-query";

// Layouts
import LayoutDefault from "./layouts/LayoutDefault";

// Views
import Home from "./views/Home";

const queryClient = new QueryClient();

const App = () => {
  const childRef = useRef();
  let location = useLocation();

  useEffect(() => {
    document.body.classList.add("is-loaded");
    childRef.current.init();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [location]);

  return (
    <ScrollReveal
      ref={childRef}
      children={() => (
        <Switch>
          <QueryClientProvider client={queryClient}>
            <AppRoute exact path="/" component={Home} layout={LayoutDefault} />
          </QueryClientProvider>
        </Switch>
      )}
    />
  );
};

export default App;
