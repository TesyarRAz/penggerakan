import React from "react";

interface CourseCardProps {
  title: string;
  description: string;
  image: string;
}

function CourseCard({ title, description, image }: CourseCardProps) {
  return (
    <div className="max-w-xs bg-white dark:bg-gray-600  rounded-lg shadow-md overflow-hidden m-4">
      <img className="w-full h-48 object-cover" src={image} alt={title} />
      <div className="p-4">
        <h3 className="text-lg font-bold mb-2 dark:text-white">{title}</h3>
        <p className="text-gray-600 text-sm mb-4 dark:text-white">
          {description}
        </p>
        <div className="flex items-center justify-between">
          <button
            className="text-gray-900 hover:text-white border border-gray-800 hover:bg-gray-900 focus:ring-4 focus:outline-none focus:ring-gray-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center me-2 mb-2 dark:border-gray-600 dark:text-gray-400 dark:hover:text-white dark:hover:bg-gray-600 dark:focus:ring-gray-800"
            onClick={() => console.log("View Course")}
          >
            View Course
          </button>
          <button
            className="text-red-700 hover:text-white border border-red-700 hover:bg-red-800 focus:ring-4 focus:outline-none focus:ring-red-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center me-2 mb-2 dark:border-red-500 dark:text-red-500 dark:hover:text-white dark:hover:bg-red-600 dark:focus:ring-red-900"
            onClick={() => {
              console.log("Bookmark");
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
