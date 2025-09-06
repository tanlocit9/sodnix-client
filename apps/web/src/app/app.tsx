import { Route, Routes, Link, Navigate } from 'react-router-dom';
import Layout from '../layout';
import Page1 from '../pages/page-1/page1';
import Page2 from '../pages/page-2/page2';
import { ROUTES } from '../contants/route';
import { QueryClientProvider } from '@tanstack/react-query';
import { queryClient } from '../lib/query-client';

export function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <Routes>
        <Route element={<Layout />}>
          <Route path={ROUTES.HOME.path} element={<div>Home</div>} />
          <Route path={ROUTES.PAGE1.path} element={<Page1></Page1>} />
          <Route path={ROUTES.PAGE2.path} element={<Page2></Page2>} />
          <Route path="*" element={<Navigate to={ROUTES.HOME.path} replace />} />
        </Route>
      </Routes>
    </QueryClientProvider>
  );
}

export default App;
