import React from 'react'

const layout = ({
    modal,
    children,
}: {
    modal: React.ReactNode,
    children: React.ReactNode,
}) => {
  return (
    <div>
      {modal}
      {children}
    </div>
  )
}

export default layout
