// AUTOGENERATED FILE
// +build !codeanalysis
// +build !syz_target syz_target,syz_os_windows,syz_arch_386

package windows

import . "github.com/google/syzkaller/prog"

func init() {
	RegisterTarget(&Target{OS: "windows", Arch: "386", Revision: revision_386, PtrSize: 4, Syscalls: syscalls_386, Resources: resources_386, Structs: structDescs_386, Consts: consts_386}, initTarget)
}

var resources_386 = []*ResourceDesc{
	{Name: "HANDLE", Type: &IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "intptr", TypeSize: 4}}}, Kind: []string{"HANDLE"}, Values: []uint64{4294967295}},
	{Name: "hFile", Type: &IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "intptr", TypeSize: 4}}}, Kind: []string{"HANDLE", "hFile"}, Values: []uint64{4294967295}},
}

var structDescs_386 = []*KeyedStruct{
	{Key: StructKey{Name: "SECURITY_ATTRIBUTES"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "SECURITY_ATTRIBUTES", TypeSize: 12}, Fields: []Type{
		&LenType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "len", FldName: "nLength", TypeSize: 4}}, Buf: "parent"},
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "lpSecurityDescriptor", TypeSize: 4, IsOptional: true}, Type: &StructType{Key: StructKey{Name: "SECURITY_DESCRIPTOR"}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "bInheritHandle", TypeSize: 4}}, Kind: 1, RangeEnd: 1},
	}}},
	{Key: StructKey{Name: "SECURITY_DESCRIPTOR"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "SECURITY_DESCRIPTOR", TypeSize: 4}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "stub", TypeSize: 4}}},
	}}},
}

var syscalls_386 = []*Syscall{
	{ID: 1, Name: "AcquireSRWLockExclusive", CallName: "AcquireSRWLockExclusive", Args: []Type{
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "SRWLock", TypeSize: 4}, Type: &IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "intptr", TypeSize: 4, ArgDir: 2}}}},
	}},
	{ID: 2, Name: "AcquireSRWLockShared", CallName: "AcquireSRWLockShared", Args: []Type{
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "SRWLock", TypeSize: 4}, Type: &IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "intptr", TypeSize: 4, ArgDir: 2}}}},
	}},
	{ID: 3, Name: "ActivateActCtx", CallName: "ActivateActCtx", Args: []Type{
		&ResourceType{TypeCommon: TypeCommon{TypeName: "HANDLE", FldName: "hActCtx", TypeSize: 4}},
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "lpCookie", TypeSize: 4}, Type: &IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int64", TypeSize: 4, ArgDir: 2}}}},
	}},
	{ID: 4, Name: "AddAtomA", CallName: "AddAtomA", Args: []Type{
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "lpString", TypeSize: 4}, Type: &IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", TypeSize: 1, ArgDir: 2}}}},
	}},
	{ID: 5, Name: "AddConsoleAliasA", CallName: "AddConsoleAliasA", Args: []Type{
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "Source", TypeSize: 4}, Type: &IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", TypeSize: 1, ArgDir: 2}}}},
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "Target", TypeSize: 4}, Type: &IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", TypeSize: 1, ArgDir: 2}}}},
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "ExeName", TypeSize: 4}, Type: &IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", TypeSize: 1, ArgDir: 2}}}},
	}},
	{ID: 6, Name: "AddRefActCtx", CallName: "AddRefActCtx", Args: []Type{
		&ResourceType{TypeCommon: TypeCommon{TypeName: "HANDLE", FldName: "hActCtx", TypeSize: 4}},
	}},
	{ID: 7, Name: "AllocateUserPhysicalPages", CallName: "AllocateUserPhysicalPages", Args: []Type{
		&ResourceType{TypeCommon: TypeCommon{TypeName: "HANDLE", FldName: "hProcess", TypeSize: 4}},
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "NumberOfPages", TypeSize: 4}, Type: &IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int64", TypeSize: 4, ArgDir: 2}}}},
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "PageArray", TypeSize: 4}, Type: &IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int64", TypeSize: 4, ArgDir: 2}}}},
	}},
	{ID: 8, Name: "AreFileApisANSI", CallName: "AreFileApisANSI"},
	{ID: 9, Name: "AssignProcessToJobObject", CallName: "AssignProcessToJobObject", Args: []Type{
		&ResourceType{TypeCommon: TypeCommon{TypeName: "HANDLE", FldName: "hJob", TypeSize: 4}},
		&ResourceType{TypeCommon: TypeCommon{TypeName: "HANDLE", FldName: "hProcess", TypeSize: 4}},
	}},
	{ID: 10, Name: "AttachConsole", CallName: "AttachConsole", Args: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "dwProcessId", TypeSize: 4}}},
	}},
	{ID: 11, Name: "BackupRead", CallName: "BackupRead", Args: []Type{
		&ResourceType{TypeCommon: TypeCommon{TypeName: "HANDLE", FldName: "hFile", TypeSize: 4}},
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "lpBuffer", TypeSize: 4}, Type: &IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", TypeSize: 1, ArgDir: 2}}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "nNumberOfBytesToRead", TypeSize: 4}}},
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "lpNumberOfBytesRead", TypeSize: 4}, Type: &IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", TypeSize: 4, ArgDir: 2}}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "bAbort", TypeSize: 4}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "bProcessSecurity", TypeSize: 4}}},
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "lpContext", TypeSize: 4}, Type: &PtrType{TypeCommon: TypeCommon{TypeName: "ptr", TypeSize: 4}, Type: &BufferType{TypeCommon: TypeCommon{TypeName: "array", ArgDir: 2}}}},
	}},
	{ID: 12, Name: "BackupSeek", CallName: "BackupSeek", Args: []Type{
		&ResourceType{TypeCommon: TypeCommon{TypeName: "HANDLE", FldName: "hFile", TypeSize: 4}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "dwLowBytesToSeek", TypeSize: 4}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "dwHighBytesToSeek", TypeSize: 4}}},
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "lpdwLowByteSeeked", TypeSize: 4}, Type: &IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", TypeSize: 4, ArgDir: 2}}}},
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "lpdwHighByteSeeked", TypeSize: 4}, Type: &IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", TypeSize: 4, ArgDir: 2}}}},
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "lpContext", TypeSize: 4}, Type: &PtrType{TypeCommon: TypeCommon{TypeName: "ptr", TypeSize: 4}, Type: &BufferType{TypeCommon: TypeCommon{TypeName: "array", ArgDir: 2}}}},
	}},
	{ID: 13, Name: "BackupWrite", CallName: "BackupWrite", Args: []Type{
		&ResourceType{TypeCommon: TypeCommon{TypeName: "HANDLE", FldName: "hFile", TypeSize: 4}},
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "lpBuffer", TypeSize: 4}, Type: &IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", TypeSize: 1, ArgDir: 2}}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "nNumberOfBytesToWrite", TypeSize: 4}}},
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "lpNumberOfBytesWritten", TypeSize: 4}, Type: &IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", TypeSize: 4, ArgDir: 2}}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "bAbort", TypeSize: 4}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "bProcessSecurity", TypeSize: 4}}},
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "lpContext", TypeSize: 4}, Type: &PtrType{TypeCommon: TypeCommon{TypeName: "ptr", TypeSize: 4}, Type: &BufferType{TypeCommon: TypeCommon{TypeName: "array", ArgDir: 2}}}},
	}},
	{ID: 14, Name: "Beep", CallName: "Beep", Args: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "dwFreq", TypeSize: 4}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "dwDuration", TypeSize: 4}}},
	}},
	{ID: 15, Name: "BeginUpdateResourceA", CallName: "BeginUpdateResourceA", Args: []Type{
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "pFileName", TypeSize: 4}, Type: &IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", TypeSize: 1, ArgDir: 2}}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "bDeleteExistingResources", TypeSize: 4}}},
	}, Ret: &ResourceType{TypeCommon: TypeCommon{TypeName: "HANDLE", FldName: "ret", TypeSize: 4, ArgDir: 1}}},
	{ID: 16, Name: "BindIoCompletionCallback", CallName: "BindIoCompletionCallback", Args: []Type{
		&ResourceType{TypeCommon: TypeCommon{TypeName: "HANDLE", FldName: "FileHandle", TypeSize: 4}},
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "Function", TypeSize: 4}, Type: &IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "intptr", TypeSize: 4, ArgDir: 2}}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "Flags", TypeSize: 4}}},
	}},
	{ID: 17, Name: "BuildCommDCBA", CallName: "BuildCommDCBA", Args: []Type{
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "lpDef", TypeSize: 4}, Type: &IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", TypeSize: 1, ArgDir: 2}}}},
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "lpDCB", TypeSize: 4}, Type: &IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "intptr", TypeSize: 4, ArgDir: 2}}}},
	}},
	{ID: 18, Name: "BuildCommDCBAndTimeoutsA", CallName: "BuildCommDCBAndTimeoutsA", Args: []Type{
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "lpDef", TypeSize: 4}, Type: &IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", TypeSize: 1, ArgDir: 2}}}},
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "lpDCB", TypeSize: 4}, Type: &IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "intptr", TypeSize: 4, ArgDir: 2}}}},
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "lpCommTimeouts", TypeSize: 4}, Type: &IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "intptr", TypeSize: 4, ArgDir: 2}}}},
	}},
	{ID: 19, Name: "CallNamedPipeA", CallName: "CallNamedPipeA", Args: []Type{
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "lpNamedPipeName", TypeSize: 4}, Type: &IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", TypeSize: 1, ArgDir: 2}}}},
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "lpInBuffer", TypeSize: 4}, Type: &BufferType{TypeCommon: TypeCommon{TypeName: "array", ArgDir: 2}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "nInBufferSize", TypeSize: 4}}},
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "lpOutBuffer", TypeSize: 4}, Type: &BufferType{TypeCommon: TypeCommon{TypeName: "array", ArgDir: 2}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "nOutBufferSize", TypeSize: 4}}},
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "lpBytesRead", TypeSize: 4}, Type: &IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", TypeSize: 4, ArgDir: 2}}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "nTimeOut", TypeSize: 4}}},
	}},
	{ID: 20, Name: "CancelDeviceWakeupRequest", CallName: "CancelDeviceWakeupRequest", Args: []Type{
		&ResourceType{TypeCommon: TypeCommon{TypeName: "HANDLE", FldName: "hDevice", TypeSize: 4}},
	}},
	{ID: 21, Name: "CancelIo", CallName: "CancelIo", Args: []Type{
		&ResourceType{TypeCommon: TypeCommon{TypeName: "HANDLE", FldName: "hFile", TypeSize: 4}},
	}},
	{ID: 22, Name: "CancelTimerQueueTimer", CallName: "CancelTimerQueueTimer", Args: []Type{
		&ResourceType{TypeCommon: TypeCommon{TypeName: "HANDLE", FldName: "TimerQueue", TypeSize: 4}},
		&ResourceType{TypeCommon: TypeCommon{TypeName: "HANDLE", FldName: "Timer", TypeSize: 4}},
	}},
	{ID: 23, Name: "CancelWaitableTimer", CallName: "CancelWaitableTimer", Args: []Type{
		&ResourceType{TypeCommon: TypeCommon{TypeName: "HANDLE", FldName: "hTimer", TypeSize: 4}},
	}},
	{ID: 24, Name: "ChangeTimerQueueTimer", CallName: "ChangeTimerQueueTimer", Args: []Type{
		&ResourceType{TypeCommon: TypeCommon{TypeName: "HANDLE", FldName: "TimerQueue", TypeSize: 4}},
		&ResourceType{TypeCommon: TypeCommon{TypeName: "HANDLE", FldName: "Timer", TypeSize: 4}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "DueTime", TypeSize: 4}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "Period", TypeSize: 4}}},
	}},
	{ID: 25, Name: "CheckNameLegalDOS8Dot3A", CallName: "CheckNameLegalDOS8Dot3A", Args: []Type{
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "lpName", TypeSize: 4}, Type: &IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", TypeSize: 1, ArgDir: 2}}}},
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "lpOemName", TypeSize: 4}, Type: &IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", TypeSize: 1, ArgDir: 2}}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "OemNameSize", TypeSize: 4}}},
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "pbNameContainsSpaces", TypeSize: 4}, Type: &IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", TypeSize: 4, ArgDir: 2}}}},
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "pbNameLegal", TypeSize: 4}, Type: &IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", TypeSize: 4, ArgDir: 2}}}},
	}},
	{ID: 26, Name: "CheckRemoteDebuggerPresent", CallName: "CheckRemoteDebuggerPresent", Args: []Type{
		&ResourceType{TypeCommon: TypeCommon{TypeName: "HANDLE", FldName: "hProcess", TypeSize: 4}},
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "pbDebuggerPresent", TypeSize: 4}, Type: &IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", TypeSize: 4, ArgDir: 2}}}},
	}},
	{ID: 27, Name: "ClearCommBreak", CallName: "ClearCommBreak", Args: []Type{
		&ResourceType{TypeCommon: TypeCommon{TypeName: "HANDLE", FldName: "hFile", TypeSize: 4}},
	}},
	{ID: 28, Name: "ClearCommError", CallName: "ClearCommError", Args: []Type{
		&ResourceType{TypeCommon: TypeCommon{TypeName: "HANDLE", FldName: "hFile", TypeSize: 4}},
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "lpErrors", TypeSize: 4}, Type: &IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", TypeSize: 4, ArgDir: 2}}}},
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "lpStat", TypeSize: 4}, Type: &IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "intptr", TypeSize: 4, ArgDir: 2}}}},
	}},
	{ID: 29, Name: "CommConfigDialogA", CallName: "CommConfigDialogA", Args: []Type{
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "lpszName", TypeSize: 4}, Type: &IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", TypeSize: 1, ArgDir: 2}}}},
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "hWnd", TypeSize: 4}, Type: &IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "intptr", TypeSize: 4, ArgDir: 2}}}},
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "lpCC", TypeSize: 4}, Type: &IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "intptr", TypeSize: 4, ArgDir: 2}}}},
	}},
	{ID: 30, Name: "CompareFileTime", CallName: "CompareFileTime", Args: []Type{
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "lpFileTime1", TypeSize: 4}, Type: &IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "intptr", TypeSize: 4, ArgDir: 2}}}},
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "lpFileTime2", TypeSize: 4}, Type: &IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "intptr", TypeSize: 4, ArgDir: 2}}}},
	}},
	{ID: 31, Name: "CloseHandle", CallName: "CloseHandle", Args: []Type{
		&ResourceType{TypeCommon: TypeCommon{TypeName: "HANDLE", FldName: "hObject", TypeSize: 4}},
	}},
	{ID: 32, Name: "CreateFileA", CallName: "CreateFileA", Args: []Type{
		&PtrType{TypeCommon{TypeName: "ptr", FldName: "lpFileName", TypeSize: 4}, &BufferType{TypeCommon: TypeCommon{TypeName: "filename", IsVarlen: true}, Kind: 3}},
		&FlagsType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "file_access_rights", FldName: "dwDesiredAccess", TypeSize: 4}}, Vals: []uint64{1, 1, 2, 2, 4, 4, 4, 8, 16, 32, 32, 64, 128, 256, 65536, 131072, 262144, 524288, 1048576, 2032127}},
		&FlagsType{IntTypeCommon{TypeCommon: TypeCommon{TypeName: "file_share_mode", FldName: "dwShareMode", TypeSize: 4}}, []uint64{1, 2, 4}, true},
		&PtrType{TypeCommon{TypeName: "ptr", FldName: "lpSecurityAttributes", TypeSize: 4, IsOptional: true}, &StructType{Key: StructKey{Name: "SECURITY_ATTRIBUTES"}}},
		&FlagsType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "file_create_disposition", FldName: "dwCreationDisposition", TypeSize: 4}}, Vals: []uint64{1, 2, 3, 4, 5}},
		&FlagsType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "file_attributes", FldName: "dwFlagsAndAttributes", TypeSize: 4}}, Vals: []uint64{0, 1, 2, 4, 32, 128, 256, 4096, 16384, 65536, 131072, 196608, 262144, 524288, 1048576, 2097152, 8388608, 16777216, 33554432, 67108864, 134217728, 268435456, 536870912, 1073741824, 2147483648}},
		&ResourceType{TypeCommon: TypeCommon{TypeName: "HANDLE", FldName: "hTemplateFile", TypeSize: 4, IsOptional: true}},
	}, Ret: &ResourceType{TypeCommon: TypeCommon{TypeName: "hFile", FldName: "ret", TypeSize: 4, ArgDir: 1}}},
	{ID: 33, Name: "VirtualAlloc", CallName: "VirtualAlloc", Args: []Type{
		&VmaType{TypeCommon: TypeCommon{TypeName: "vma", FldName: "lpAddress", TypeSize: 4}},
		&LenType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "len", FldName: "dwSize", TypeSize: 4}}, Buf: "lpAddress"},
		&FlagsType{IntTypeCommon{TypeCommon: TypeCommon{TypeName: "allocation_type", FldName: "flAllocationType", TypeSize: 4}}, []uint64{4096, 8192, 524288, 1048576, 2097152, 4194304, 16777216, 536870912}, true},
		&FlagsType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "protect_flags", FldName: "flProtect", TypeSize: 4}}, Vals: []uint64{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 536870912, 1073741824, 1073741824, 2147483648}},
	}},
	{ID: 34, Name: "GetProcessId", CallName: "GetProcessId", Args: []Type{
		&ResourceType{TypeCommon: TypeCommon{TypeName: "HANDLE", FldName: "Process", TypeSize: 4}},
	}},
	{ID: 35, Name: "syz_execute_func", CallName: "syz_execute_func", Args: []Type{
		&PtrType{TypeCommon{TypeName: "ptr", FldName: "text", TypeSize: 4}, &BufferType{TypeCommon: TypeCommon{TypeName: "text", IsVarlen: true}, Kind: 4}},
	}},
}

var consts_386 = []ConstValue{
	{Name: "CREATE_ALWAYS", Value: 2},
	{Name: "CREATE_NEW", Value: 1},
	{Name: "DELETE", Value: 65536},
	{Name: "FILE_ADD_FILE", Value: 2},
	{Name: "FILE_ADD_SUBDIRECTORY", Value: 4},
	{Name: "FILE_ALL_ACCESS", Value: 2032127},
	{Name: "FILE_APPEND_DATA", Value: 4},
	{Name: "FILE_ATTRIBUTE_ARCHIVE", Value: 32},
	{Name: "FILE_ATTRIBUTE_ENCRYPTED", Value: 16384},
	{Name: "FILE_ATTRIBUTE_HIDDEN", Value: 2},
	{Name: "FILE_ATTRIBUTE_NORMAL", Value: 128},
	{Name: "FILE_ATTRIBUTE_OFFLINE", Value: 4096},
	{Name: "FILE_ATTRIBUTE_READONLY", Value: 1},
	{Name: "FILE_ATTRIBUTE_SYSTEM", Value: 4},
	{Name: "FILE_ATTRIBUTE_TEMPORARY", Value: 256},
	{Name: "FILE_CREATE_PIPE_INSTANCE", Value: 4},
	{Name: "FILE_DELETE_CHILD", Value: 64},
	{Name: "FILE_EXECUTE", Value: 32},
	{Name: "FILE_FLAG_BACKUP_SEMANTICS", Value: 33554432},
	{Name: "FILE_FLAG_DELETE_ON_CLOSE", Value: 67108864},
	{Name: "FILE_FLAG_NO_BUFFERING", Value: 536870912},
	{Name: "FILE_FLAG_OPEN_NO_RECALL", Value: 1048576},
	{Name: "FILE_FLAG_OPEN_REPARSE_POINT", Value: 2097152},
	{Name: "FILE_FLAG_OVERLAPPED", Value: 1073741824},
	{Name: "FILE_FLAG_POSIX_SEMANTICS", Value: 16777216},
	{Name: "FILE_FLAG_RANDOM_ACCESS", Value: 268435456},
	{Name: "FILE_FLAG_SEQUENTIAL_SCAN", Value: 134217728},
	{Name: "FILE_FLAG_SESSION_AWARE", Value: 8388608},
	{Name: "FILE_FLAG_WRITE_THROUGH", Value: 2147483648},
	{Name: "FILE_LIST_DIRECTORY", Value: 1},
	{Name: "FILE_READ_ATTRIBUTES", Value: 128},
	{Name: "FILE_READ_DATA", Value: 1},
	{Name: "FILE_READ_EA", Value: 8},
	{Name: "FILE_SHARE_DELETE", Value: 4},
	{Name: "FILE_SHARE_READ", Value: 1},
	{Name: "FILE_SHARE_WRITE", Value: 2},
	{Name: "FILE_TRAVERSE", Value: 32},
	{Name: "FILE_WRITE_ATTRIBUTES", Value: 256},
	{Name: "FILE_WRITE_DATA", Value: 2},
	{Name: "FILE_WRITE_EA", Value: 16},
	{Name: "INVALID_HANDLE_VALUE", Value: 4294967295},
	{Name: "MEM_COMMIT", Value: 4096},
	{Name: "MEM_LARGE_PAGES", Value: 536870912},
	{Name: "MEM_PHYSICAL", Value: 4194304},
	{Name: "MEM_RESERVE", Value: 8192},
	{Name: "MEM_RESET", Value: 524288},
	{Name: "MEM_RESET_UNDO", Value: 16777216},
	{Name: "MEM_TOP_DOWN", Value: 1048576},
	{Name: "MEM_WRITE_WATCH", Value: 2097152},
	{Name: "OPEN_ALWAYS", Value: 4},
	{Name: "OPEN_EXISTING", Value: 3},
	{Name: "PAGE_ENCLAVE_THREAD_CONTROL", Value: 2147483648},
	{Name: "PAGE_ENCLAVE_UNVALIDATED", Value: 536870912},
	{Name: "PAGE_EXECUTE", Value: 16},
	{Name: "PAGE_EXECUTE_READ", Value: 32},
	{Name: "PAGE_EXECUTE_READWRITE", Value: 64},
	{Name: "PAGE_EXECUTE_WRITECOPY", Value: 128},
	{Name: "PAGE_GUARD", Value: 256},
	{Name: "PAGE_NOACCESS", Value: 1},
	{Name: "PAGE_NOCACHE", Value: 512},
	{Name: "PAGE_READONLY", Value: 2},
	{Name: "PAGE_READWRITE", Value: 4},
	{Name: "PAGE_TARGETS_INVALID", Value: 1073741824},
	{Name: "PAGE_TARGETS_NO_UPDATE", Value: 1073741824},
	{Name: "PAGE_WRITECOMBINE", Value: 1024},
	{Name: "PAGE_WRITECOPY", Value: 8},
	{Name: "READ_CONTROL", Value: 131072},
	{Name: "SECURITY_ANONYMOUS"},
	{Name: "SECURITY_CONTEXT_TRACKING", Value: 262144},
	{Name: "SECURITY_DELEGATION", Value: 196608},
	{Name: "SECURITY_EFFECTIVE_ONLY", Value: 524288},
	{Name: "SECURITY_IDENTIFICATION", Value: 65536},
	{Name: "SECURITY_IMPERSONATION", Value: 131072},
	{Name: "SYNCHRONIZE", Value: 1048576},
	{Name: "TRUNCATE_EXISTING", Value: 5},
	{Name: "WRITE_DAC", Value: 262144},
	{Name: "WRITE_OWNER", Value: 524288},
}

const revision_386 = "97353516850eeffbb71bb94a12f1991ff6ad15c4"
