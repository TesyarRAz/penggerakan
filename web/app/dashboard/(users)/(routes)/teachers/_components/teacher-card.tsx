"use client"
import { useRouter } from 'next/navigation'
import React from 'react'
import deleteUser from '../../users/_actions/delete-user-action'
import { FaTrashAlt } from 'react-icons/fa'
import { TbListDetails } from 'react-icons/tb'
import { Session } from 'next-auth'
import deleteTeacher from '../_actions/delete-teacher-action'

interface TeacherCardProps {
  session: Session
  teacher: TeacherResponse
}

const TeacherCard = ({
  session,
  teacher: {
    id,
    user_id,
    name,
    profile_image,
  }
}: TeacherCardProps) => {
  const router = useRouter()

  const handleDelete = async () => {
    if (!session) return

    if (!confirm('Are you sure you want to delete this course?')) return

    const ok = await deleteTeacher(session, id)

    if (ok) {
      router.refresh()
    } else {
      alert('Failed to delete course')
    }
  }

  return (
    <div className="max-w-xs bg-white dark:bg-gray-600  rounded-lg shadow-md overflow-hidden m-4">
      <div className="bg-blue-300 py-2 pl-11">
        <picture>
          <img
            src={profile_image}
            alt={name}
            className="w-48 h-48 rounded-full object-fill"
          />
        </picture>
      </div>
      <div className="p-4">
        <h3 className="text-lg font-bold mb-2 dark:text-white">
          {name.toLocaleUpperCase()}
        </h3>
        <p className="text-gray-600 text-sm  dark:text-white">
          ID : <span className="text-xs">{id}</span>
        </p>
        <p className="text-gray-600 text-sm  dark:text-white">
          UserID : <span className="text-xs">{user_id}</span>
        </p>
        <div className="flex items-center justify-between mt-3">
          <button
            className="text-blue-700 hover:text-white border border-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center me-2 mb-2 dark:border-blue-500 dark:text-blue-500 dark:hover:text-white dark:hover:bg-blue-500 dark:focus:ring-blue-800"
            onClick={() => console.log("Details")}
          >
            <div className="flex justify-center items-center">
              <TbListDetails className="mr-1" />
              <span>Details</span>
            </div>
          </button>
          <button
            type="button"
            className="text-red-700 hover:text-white border border-red-700 hover:bg-red-800 focus:ring-4 focus:outline-none focus:ring-red-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center me-2 mb-2 dark:border-red-500 dark:text-red-500 dark:hover:text-white dark:hover:bg-red-600 dark:focus:ring-red-900"
          >
            <div className="flex items-center justify-center">
              <FaTrashAlt className="mr-1" />
              <span>Delete</span>
            </div>
          </button>
        </div>
      </div>
    </div>
  );
}

export default TeacherCard
