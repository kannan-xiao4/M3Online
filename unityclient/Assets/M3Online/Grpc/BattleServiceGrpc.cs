// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: service/battle_service.proto
// </auto-generated>
#pragma warning disable 0414, 1591
#region Designer generated code

using grpc = global::Grpc.Core;

namespace M3Online.Grpc.BattleService {
  public static partial class BattleService
  {
    static readonly string __ServiceName = "BattleService";

    static readonly grpc::Marshaller<global::M3Online.Grpc.BattleService.EnterRequest> __Marshaller_EnterRequest = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::M3Online.Grpc.BattleService.EnterRequest.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::M3Online.Grpc.BattleService.EnterResponse> __Marshaller_EnterResponse = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::M3Online.Grpc.BattleService.EnterResponse.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::M3Online.Grpc.BattleService.ConnectionRequest> __Marshaller_ConnectionRequest = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::M3Online.Grpc.BattleService.ConnectionRequest.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::M3Online.Grpc.BattleService.EnemySituation> __Marshaller_EnemySituation = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::M3Online.Grpc.BattleService.EnemySituation.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::M3Online.Grpc.BattleService.AttackRequest> __Marshaller_AttackRequest = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::M3Online.Grpc.BattleService.AttackRequest.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::M3Online.Grpc.BattleService.SessionSummary> __Marshaller_SessionSummary = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::M3Online.Grpc.BattleService.SessionSummary.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::M3Online.Grpc.BattleService.UserList> __Marshaller_UserList = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::M3Online.Grpc.BattleService.UserList.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::M3Online.Grpc.BattleService.ExitRequest> __Marshaller_ExitRequest = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::M3Online.Grpc.BattleService.ExitRequest.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::M3Online.Grpc.BattleService.ExitResponse> __Marshaller_ExitResponse = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::M3Online.Grpc.BattleService.ExitResponse.Parser.ParseFrom);

    static readonly grpc::Method<global::M3Online.Grpc.BattleService.EnterRequest, global::M3Online.Grpc.BattleService.EnterResponse> __Method_EnterBattle = new grpc::Method<global::M3Online.Grpc.BattleService.EnterRequest, global::M3Online.Grpc.BattleService.EnterResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "EnterBattle",
        __Marshaller_EnterRequest,
        __Marshaller_EnterResponse);

    static readonly grpc::Method<global::M3Online.Grpc.BattleService.ConnectionRequest, global::M3Online.Grpc.BattleService.EnemySituation> __Method_Connect = new grpc::Method<global::M3Online.Grpc.BattleService.ConnectionRequest, global::M3Online.Grpc.BattleService.EnemySituation>(
        grpc::MethodType.ServerStreaming,
        __ServiceName,
        "Connect",
        __Marshaller_ConnectionRequest,
        __Marshaller_EnemySituation);

    static readonly grpc::Method<global::M3Online.Grpc.BattleService.AttackRequest, global::M3Online.Grpc.BattleService.SessionSummary> __Method_Attack = new grpc::Method<global::M3Online.Grpc.BattleService.AttackRequest, global::M3Online.Grpc.BattleService.SessionSummary>(
        grpc::MethodType.ClientStreaming,
        __ServiceName,
        "Attack",
        __Marshaller_AttackRequest,
        __Marshaller_SessionSummary);

    static readonly grpc::Method<global::M3Online.Grpc.BattleService.ConnectionRequest, global::M3Online.Grpc.BattleService.UserList> __Method_ReceiveUsers = new grpc::Method<global::M3Online.Grpc.BattleService.ConnectionRequest, global::M3Online.Grpc.BattleService.UserList>(
        grpc::MethodType.ServerStreaming,
        __ServiceName,
        "ReceiveUsers",
        __Marshaller_ConnectionRequest,
        __Marshaller_UserList);

    static readonly grpc::Method<global::M3Online.Grpc.BattleService.ExitRequest, global::M3Online.Grpc.BattleService.ExitResponse> __Method_Exit = new grpc::Method<global::M3Online.Grpc.BattleService.ExitRequest, global::M3Online.Grpc.BattleService.ExitResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "Exit",
        __Marshaller_ExitRequest,
        __Marshaller_ExitResponse);

    /// <summary>Service descriptor</summary>
    public static global::Google.Protobuf.Reflection.ServiceDescriptor Descriptor
    {
      get { return global::M3Online.Grpc.BattleService.BattleServiceReflection.Descriptor.Services[0]; }
    }

    /// <summary>Base class for server-side implementations of BattleService</summary>
    [grpc::BindServiceMethod(typeof(BattleService), "BindService")]
    public abstract partial class BattleServiceBase
    {
      public virtual global::System.Threading.Tasks.Task<global::M3Online.Grpc.BattleService.EnterResponse> EnterBattle(global::M3Online.Grpc.BattleService.EnterRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      public virtual global::System.Threading.Tasks.Task Connect(global::M3Online.Grpc.BattleService.ConnectionRequest request, grpc::IServerStreamWriter<global::M3Online.Grpc.BattleService.EnemySituation> responseStream, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      public virtual global::System.Threading.Tasks.Task<global::M3Online.Grpc.BattleService.SessionSummary> Attack(grpc::IAsyncStreamReader<global::M3Online.Grpc.BattleService.AttackRequest> requestStream, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      public virtual global::System.Threading.Tasks.Task ReceiveUsers(global::M3Online.Grpc.BattleService.ConnectionRequest request, grpc::IServerStreamWriter<global::M3Online.Grpc.BattleService.UserList> responseStream, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      public virtual global::System.Threading.Tasks.Task<global::M3Online.Grpc.BattleService.ExitResponse> Exit(global::M3Online.Grpc.BattleService.ExitRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

    }

    /// <summary>Client for BattleService</summary>
    public partial class BattleServiceClient : grpc::ClientBase<BattleServiceClient>
    {
      /// <summary>Creates a new client for BattleService</summary>
      /// <param name="channel">The channel to use to make remote calls.</param>
      public BattleServiceClient(grpc::Channel channel) : base(channel)
      {
      }
      /// <summary>Creates a new client for BattleService that uses a custom <c>CallInvoker</c>.</summary>
      /// <param name="callInvoker">The callInvoker to use to make remote calls.</param>
      public BattleServiceClient(grpc::CallInvoker callInvoker) : base(callInvoker)
      {
      }
      /// <summary>Protected parameterless constructor to allow creation of test doubles.</summary>
      protected BattleServiceClient() : base()
      {
      }
      /// <summary>Protected constructor to allow creation of configured clients.</summary>
      /// <param name="configuration">The client configuration.</param>
      protected BattleServiceClient(ClientBaseConfiguration configuration) : base(configuration)
      {
      }

      public virtual global::M3Online.Grpc.BattleService.EnterResponse EnterBattle(global::M3Online.Grpc.BattleService.EnterRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return EnterBattle(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual global::M3Online.Grpc.BattleService.EnterResponse EnterBattle(global::M3Online.Grpc.BattleService.EnterRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_EnterBattle, null, options, request);
      }
      public virtual grpc::AsyncUnaryCall<global::M3Online.Grpc.BattleService.EnterResponse> EnterBattleAsync(global::M3Online.Grpc.BattleService.EnterRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return EnterBattleAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncUnaryCall<global::M3Online.Grpc.BattleService.EnterResponse> EnterBattleAsync(global::M3Online.Grpc.BattleService.EnterRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_EnterBattle, null, options, request);
      }
      public virtual grpc::AsyncServerStreamingCall<global::M3Online.Grpc.BattleService.EnemySituation> Connect(global::M3Online.Grpc.BattleService.ConnectionRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return Connect(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncServerStreamingCall<global::M3Online.Grpc.BattleService.EnemySituation> Connect(global::M3Online.Grpc.BattleService.ConnectionRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncServerStreamingCall(__Method_Connect, null, options, request);
      }
      public virtual grpc::AsyncClientStreamingCall<global::M3Online.Grpc.BattleService.AttackRequest, global::M3Online.Grpc.BattleService.SessionSummary> Attack(grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return Attack(new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncClientStreamingCall<global::M3Online.Grpc.BattleService.AttackRequest, global::M3Online.Grpc.BattleService.SessionSummary> Attack(grpc::CallOptions options)
      {
        return CallInvoker.AsyncClientStreamingCall(__Method_Attack, null, options);
      }
      public virtual grpc::AsyncServerStreamingCall<global::M3Online.Grpc.BattleService.UserList> ReceiveUsers(global::M3Online.Grpc.BattleService.ConnectionRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return ReceiveUsers(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncServerStreamingCall<global::M3Online.Grpc.BattleService.UserList> ReceiveUsers(global::M3Online.Grpc.BattleService.ConnectionRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncServerStreamingCall(__Method_ReceiveUsers, null, options, request);
      }
      public virtual global::M3Online.Grpc.BattleService.ExitResponse Exit(global::M3Online.Grpc.BattleService.ExitRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return Exit(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual global::M3Online.Grpc.BattleService.ExitResponse Exit(global::M3Online.Grpc.BattleService.ExitRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_Exit, null, options, request);
      }
      public virtual grpc::AsyncUnaryCall<global::M3Online.Grpc.BattleService.ExitResponse> ExitAsync(global::M3Online.Grpc.BattleService.ExitRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return ExitAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncUnaryCall<global::M3Online.Grpc.BattleService.ExitResponse> ExitAsync(global::M3Online.Grpc.BattleService.ExitRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_Exit, null, options, request);
      }
      /// <summary>Creates a new instance of client from given <c>ClientBaseConfiguration</c>.</summary>
      protected override BattleServiceClient NewInstance(ClientBaseConfiguration configuration)
      {
        return new BattleServiceClient(configuration);
      }
    }

    /// <summary>Creates service definition that can be registered with a server</summary>
    /// <param name="serviceImpl">An object implementing the server-side handling logic.</param>
    public static grpc::ServerServiceDefinition BindService(BattleServiceBase serviceImpl)
    {
      return grpc::ServerServiceDefinition.CreateBuilder()
          .AddMethod(__Method_EnterBattle, serviceImpl.EnterBattle)
          .AddMethod(__Method_Connect, serviceImpl.Connect)
          .AddMethod(__Method_Attack, serviceImpl.Attack)
          .AddMethod(__Method_ReceiveUsers, serviceImpl.ReceiveUsers)
          .AddMethod(__Method_Exit, serviceImpl.Exit).Build();
    }

    /// <summary>Register service method with a service binder with or without implementation. Useful when customizing the  service binding logic.
    /// Note: this method is part of an experimental API that can change or be removed without any prior notice.</summary>
    /// <param name="serviceBinder">Service methods will be bound by calling <c>AddMethod</c> on this object.</param>
    /// <param name="serviceImpl">An object implementing the server-side handling logic.</param>
    public static void BindService(grpc::ServiceBinderBase serviceBinder, BattleServiceBase serviceImpl)
    {
      serviceBinder.AddMethod(__Method_EnterBattle, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::M3Online.Grpc.BattleService.EnterRequest, global::M3Online.Grpc.BattleService.EnterResponse>(serviceImpl.EnterBattle));
      serviceBinder.AddMethod(__Method_Connect, serviceImpl == null ? null : new grpc::ServerStreamingServerMethod<global::M3Online.Grpc.BattleService.ConnectionRequest, global::M3Online.Grpc.BattleService.EnemySituation>(serviceImpl.Connect));
      serviceBinder.AddMethod(__Method_Attack, serviceImpl == null ? null : new grpc::ClientStreamingServerMethod<global::M3Online.Grpc.BattleService.AttackRequest, global::M3Online.Grpc.BattleService.SessionSummary>(serviceImpl.Attack));
      serviceBinder.AddMethod(__Method_ReceiveUsers, serviceImpl == null ? null : new grpc::ServerStreamingServerMethod<global::M3Online.Grpc.BattleService.ConnectionRequest, global::M3Online.Grpc.BattleService.UserList>(serviceImpl.ReceiveUsers));
      serviceBinder.AddMethod(__Method_Exit, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::M3Online.Grpc.BattleService.ExitRequest, global::M3Online.Grpc.BattleService.ExitResponse>(serviceImpl.Exit));
    }

  }
}
#endregion
