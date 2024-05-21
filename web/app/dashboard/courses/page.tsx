import { authOptions } from "@/app/api/auth/[...nextauth]/route";
import axios from "axios";
import { getServerSession } from "next-auth";
import React from "react";
import { getCourses } from "./_actions/get-courses-action";
import { redirect } from "next/navigation";
import CourseCard from "./_components/course-card";
import CourseList from "./_components/course-list";

const CoursePage = async () => {
  const session = await getServerSession(authOptions);

  if (!session) {
    return redirect("/auth/signin?callback=/dashboard/courses");
  }

  const courses = await getCourses(session);

  return (
    <CourseList courses={courses.data} />
  );
};

export default CoursePage;
