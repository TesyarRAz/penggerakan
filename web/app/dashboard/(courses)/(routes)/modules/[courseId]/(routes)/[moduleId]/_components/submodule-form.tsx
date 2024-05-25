"use client"
import React, { useEffect, useState } from 'react'
import { Session } from "next-auth";
import { Button } from "@/components/ui/button";
import { Popover, PopoverContent, PopoverTrigger } from '@/components/ui/popover';
import { Command, CommandEmpty, CommandGroup, CommandInput, CommandItem, CommandList } from '@/components/ui/command';
import StructureModal from './structure-modal';
import { SubModuleStructureType } from '@/types/enums';
import { v4 as uuid } from 'uuid';
import { Collapsible, CollapsibleContent, CollapsibleTrigger } from '@/components/ui/collapsible';
import { FaSort, FaSortDown, FaSortUp } from 'react-icons/fa';
import { cn } from '@/lib/utils';
import PreviewStructureModal from './preview-structure-modal';

const AddItemButton = ({
    addItem
}: {
    addItem: (type: string) => void
}) => {
    const [open, setOpen] = useState(false)

    return (
        <Popover open={open} onOpenChange={setOpen}>
            <PopoverTrigger asChild>
                <Button variant="outline" className="w-[150px] justify-start">
                    AddItem
                </Button>
            </PopoverTrigger>
            <PopoverContent className="p-0" side="right" align="start">
                <Command>
                    <CommandInput placeholder="Add item..." />
                    <CommandList>
                        <CommandEmpty>No results found.</CommandEmpty>
                        <CommandGroup>
                            {Object.values(SubModuleStructureType).map((value) => (
                                <CommandItem
                                    key={value}
                                    value={value}
                                    onSelect={(value) => {
                                        addItem(value)
                                        setOpen(false)
                                    }}
                                >
                                    {value}
                                </CommandItem>
                            ))}
                        </CommandGroup>
                    </CommandList>
                </Command>
            </PopoverContent>
        </Popover>
    )
}

const ListItem = ({
    structure,
    setSelectedStructure,
    updateItems,
    deleteItem,
    ...props
}: React.HTMLAttributes<HTMLDivElement> & {
    structure: SubModuleStructure,
    setSelectedStructure: (structure: SubModuleStructure) => void,
    updateItems: () => void,
    deleteItem: () => void,
}) => {
    const [open, setOpen] = useState(true)

    return (
        <Collapsible open={open} onOpenChange={setOpen}>
            <div className="ml-2" {...props}>
                <div className="flex justify-between">
                    <div>{structure.label}</div>
                    <div>
                        <CollapsibleTrigger asChild>
                            <Button variant="ghost">
                                <FaSortUp className={cn("h-4 w-4", open ? "hidden" : "")} />
                                <FaSortDown className={cn("h-4 w-4", !open ? "hidden" : "")} />
                            </Button>
                        </CollapsibleTrigger>
                        <Button
                            variant="outline"
                            onClick={() => {
                                setSelectedStructure(structure)
                            }}
                        >
                            Edit
                        </Button>
                        <Button
                            variant="outline"
                            onClick={deleteItem}
                        >
                            Delete
                        </Button>
                    </div>
                </div>

                <div className="mb-3">
                    <CollapsibleContent>
                        {structure.children?.map((item, index) => (
                            <ListItem
                                key={index}
                                structure={item}
                                updateItems={updateItems}
                                deleteItem={() => {
                                    structure.children = structure.children?.filter((_, i) => i !== index)
                                    updateItems()
                                }}
                                setSelectedStructure={setSelectedStructure} />
                        ))}
                    </CollapsibleContent>

                    <AddItemButton addItem={(type) => {
                        structure.children = [
                            ...structure.children ?? [],
                            {
                                resource_id: uuid(),
                                resource_type: type,
                                label: "New Item",
                                value: ""
                            }
                        ]

                        updateItems()
                    }} />
                </div>
            </div>
        </Collapsible>
    )
}

const SubModuleForm = ({
    session,
    course,
    mod,
}: {
    session: Session,
    course: CourseResponse,
    mod: ModuleResponse,
}) => {
    const [modalOpen, setModalOpen] = useState(false)
    const [structureModalOpen, setStructureModalOpen] = useState(false)

    const [selectedStructure, setSelectedStructure] = useState<SubModuleStructure | null>(null)
    const [items, setItems] = useState<SubModuleStructure[]>([])

    useEffect(() => {
        setModalOpen(selectedStructure !== null)
    }, [selectedStructure])

    const addItem = (type: string) => {
        setItems([...items, {
            resource_id: uuid(),
            resource_type: type,
            label: "New Item",
            value: ""
        }])
    }

    const handleModalSubmit = (label: string, value: string) => {
        if (selectedStructure) {
            selectedStructure.label = label
            selectedStructure.value = value
        }
        setSelectedStructure(null)
        setItems([...items])
    }

    return (
        <>
            <StructureModal open={modalOpen} onOpenChange={setModalOpen} structure={selectedStructure} onSubmit={handleModalSubmit} />
            <PreviewStructureModal open={structureModalOpen} onOpenChange={setStructureModalOpen} structures={items} />
            <Button onClick={() => setStructureModalOpen(true)}>Preview</Button>
            <div className="mt-5">
                {items.map((item, index) => (
                    <ListItem
                        key={index}
                        structure={item}
                        updateItems={() => setItems([...items])}
                        deleteItem={() => setItems(items.filter((_, i) => i !== index))}
                        setSelectedStructure={setSelectedStructure} />
                ))}

                <div className="mt-5">
                    <AddItemButton addItem={addItem} />
                </div>
            </div>
        </>
    )
}

export default SubModuleForm
