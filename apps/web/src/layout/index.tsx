import React from 'react';
import { Outlet } from 'react-router-dom';
import { Sidebar } from '../components/sidebar';

const Layout: React.FC = () => {
  return (
    <div className="flex h-screen">
      <Sidebar />

      {/* Main Content */}
      <div className="flex flex-col flex-1">
        {/* Top Menu */}
        <header className="h-15 bg-blue-600 text-white flex items-center px-4 shadow-md">
          <h1 className="text-xl font-bold">Top Menu</h1>
        </header>

        {/* Content Area */}
        <main className="flex-1 p-4 bg-white">
          <Outlet />
        </main>
      </div>
    </div>
  );
};

export default Layout;
