import axios from "axios";
import React from "react";
import { getCourses } from "./_actions/get-courses-action";
import { redirect } from "next/navigation";
import CourseCard from "./_components/course-card";
import Link from "next/link";
import { FaPlus, FaSearch } from "react-icons/fa";
import { RiEqualizerFill } from "react-icons/ri";
import { Button } from "@/components/ui/button";
import { IoAddCircleOutline } from "react-icons/io5";
import BrowseLayout from "@/components/layouts/browse-layout";
import { auth } from "@/auth";

const CoursePage = async () => {
  const session = await auth()

  if (!session) {
    return redirect("/auth/signin?callback=/dashboard/courses");
  }

  const courses = await getCourses(session);

  return (
    <BrowseLayout
      title="Course List"
      tools={(
        <Button
          variant="blank"
          className="text-green-700 hover:text-white border border-green-700 hover:bg-green-800 focus:ring-4 focus:outline-none focus:ring-green-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center me-2 mb-2 dark:border-green-500 dark:text-green-500 dark:hover:text-white dark:hover:bg-green-600 dark:focus:ring-green-800"
          asChild
        >
          <Link href="/dashboard/courses/create">
            <IoAddCircleOutline className="mr-1 h-5 w-5" />
            <span>Add Course</span>
          </Link>
        </Button>
      )}
    >
      <div className="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-4 mt-5">
        {courses.data.map((item) => (
          <CourseCard
            session={session}
            key={item.id}
            course={item}
          />
        ))}
      </div>
    </BrowseLayout>
  );
};

export default CoursePage;
