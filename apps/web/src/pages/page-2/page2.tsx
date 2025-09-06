import React from 'react';
import Skeleton from '../../components/skeleton';
import { useLogic } from './useLogic';

const Page2: React.FC = () => {
  const { user, isLoading } = useLogic();

  if (isLoading) {
    return <Skeleton />;
  }

  return (
    <div className="text-center p-8">
      <h1 className="text-2xl font-bold">Welcome to Page 2</h1>
      <p className="text-gray-600">{user?.name}</p>
    </div>
  );
};

export default Page2;
