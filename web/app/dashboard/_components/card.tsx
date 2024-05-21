import Image from "next/image";
import React from "react";

interface CourseCardProps {
  title: string;
  description: string;
  image: string;
}

function CourseCard({ title, description, image }: CourseCardProps) {
  return (
    <div className="max-w-xs bg-white rounded-lg shadow-md overflow-hidden m-4">
      <div className="h-48 w-full relative">
      <Image className="object-cover" src={image} alt={title} layout="fill"/>
      </div>
      <div className="p-4">
        <h3 className="text-lg font-bold mb-2">{title}</h3>
        <p className="text-gray-600 text-sm mb-4">{description}</p>
        <div className="flex items-center justify-between">
          <button className="py-2.5 px-5 me-2 mb-2 text-sm font-medium text-gray-900 focus:outline-none bg-white rounded-lg border border-gray-200 hover:bg-gray-100 hover:text-blue-700 focus:z-10 focus:ring-4 focus:ring-gray-100 dark:focus:ring-gray-700 dark:bg-gray-800 dark:text-gray-400 dark:border-gray-600 dark:hover:text-white dark:hover:bg-gray-700">
            View Course
          </button>
          <button
            className="focus:outline-none text-white bg-red-700 hover:bg-red-800 focus:ring-4 focus:ring-red-300 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2 dark:bg-red-600 dark:hover:bg-red-700 dark:focus:ring-red-900"
            onClick={() => {
              console.log(image);
            }}
          >
            Bookmark
          </button>
        </div>
      </div>
    </div>
  );
}

export default CourseCard;
