package grpc_server_test

// func TestServerCreateLaptop(t *testing.T) {
// 	t.Parallel()

// 	laptopNoID := generator.NewLaptop()
// 	laptopNoID.Id = ""

// 	laptopInvalidID := generator.NewLaptop()
// 	laptopInvalidID.Id = "invalid-uuid"

// 	laptopDuplicateID := generator.NewLaptop()
// 	storeDuplicateID := service.NewInMemoryLaptopStore()
// 	err := storeDuplicateID.Save(laptopDuplicateID)
// 	require.Nil(t, err)

// 	testCases := []struct {
// 		name       string
// 		laptop     *pb.Laptop
// 		store      service.LaptopStore
// 		imageStore service.ImageStore
// 		code       codes.Code
// 	}{
// 		{
// 			name:       "success_with_id",
// 			laptop:     generator.NewLaptop(),
// 			store:      service.NewInMemoryLaptopStore(),
// 			imageStore: service.NewDiskImageStore("image"),
// 			code:       codes.OK,
// 		},
// 		{
// 			name:       "success_no_id",
// 			laptop:     laptopNoID,
// 			store:      service.NewInMemoryLaptopStore(),
// 			imageStore: service.NewDiskImageStore("image"),
// 			code:       codes.OK,
// 		},
// 		{
// 			name:       "failure_invalid_id",
// 			laptop:     laptopInvalidID,
// 			store:      service.NewInMemoryLaptopStore(),
// 			imageStore: service.NewDiskImageStore("image"),
// 			code:       codes.InvalidArgument,
// 		},
// 		{
// 			name:       "failure_duplicate_id",
// 			laptop:     laptopDuplicateID,
// 			store:      storeDuplicateID,
// 			imageStore: service.NewDiskImageStore("image"),
// 			code:       codes.AlreadyExists,
// 		},
// 	}

// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			t.Parallel()

// 			req := &pb.CreateLaptopRequest{
// 				Laptop: tc.laptop,
// 			}

// 			server := service.NewLaptopServer(tc.store, tc.imageStore)
// 			res, err := server.CreateLaptop(context.Background(), req)
// 			if tc.code == codes.OK {
// 				require.NoError(t, err)
// 				require.NotNil(t, res)
// 				require.NotEmpty(t, res.Id)
// 				if len(tc.laptop.Id) > 0 {
// 					require.Equal(t, tc.laptop.Id, res.Id)
// 				}
// 			} else {
// 				require.Error(t, err)
// 				require.Nil(t, res)
// 				st, ok := status.FromError(err)
// 				require.True(t, ok)
// 				require.Equal(t, tc.code, st.Code())
// 			}
// 		})
// 	}
// }
