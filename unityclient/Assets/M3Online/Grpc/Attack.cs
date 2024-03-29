// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: domain/attack.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace M3Online.Domain.Attack {

  /// <summary>Holder for reflection information generated from domain/attack.proto</summary>
  public static partial class AttackReflection {

    #region Descriptor
    /// <summary>File descriptor for domain/attack.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static AttackReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "ChNkb21haW4vYXR0YWNrLnByb3RvGhJkb21haW4vZW50ZXIucHJvdG8aEmRv",
            "bWFpbi9lbmVteS5wcm90byI9CgZBdHRhY2sSFQoFZW50ZXIYASABKAsyBi5F",
            "bnRlchIcCgxhdHRhY2tfZW5lbXkYAiABKAsyBi5FbmVteUInWgxtM29ubGlu",
            "ZS9ycGOqAhZNM09ubGluZS5Eb21haW4uQXR0YWNrYgZwcm90bzM="));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::M3Online.Domain.Enter.EnterReflection.Descriptor, global::M3Online.Domain.Enemy.EnemyReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::M3Online.Domain.Attack.Attack), global::M3Online.Domain.Attack.Attack.Parser, new[]{ "Enter", "AttackEnemy" }, null, null, null)
          }));
    }
    #endregion

  }
  #region Messages
  public sealed partial class Attack : pb::IMessage<Attack> {
    private static readonly pb::MessageParser<Attack> _parser = new pb::MessageParser<Attack>(() => new Attack());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<Attack> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::M3Online.Domain.Attack.AttackReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public Attack() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public Attack(Attack other) : this() {
      enter_ = other.enter_ != null ? other.enter_.Clone() : null;
      attackEnemy_ = other.attackEnemy_ != null ? other.attackEnemy_.Clone() : null;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public Attack Clone() {
      return new Attack(this);
    }

    /// <summary>Field number for the "enter" field.</summary>
    public const int EnterFieldNumber = 1;
    private global::M3Online.Domain.Enter.Enter enter_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::M3Online.Domain.Enter.Enter Enter {
      get { return enter_; }
      set {
        enter_ = value;
      }
    }

    /// <summary>Field number for the "attack_enemy" field.</summary>
    public const int AttackEnemyFieldNumber = 2;
    private global::M3Online.Domain.Enemy.Enemy attackEnemy_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::M3Online.Domain.Enemy.Enemy AttackEnemy {
      get { return attackEnemy_; }
      set {
        attackEnemy_ = value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as Attack);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(Attack other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (!object.Equals(Enter, other.Enter)) return false;
      if (!object.Equals(AttackEnemy, other.AttackEnemy)) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (enter_ != null) hash ^= Enter.GetHashCode();
      if (attackEnemy_ != null) hash ^= AttackEnemy.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
      if (enter_ != null) {
        output.WriteRawTag(10);
        output.WriteMessage(Enter);
      }
      if (attackEnemy_ != null) {
        output.WriteRawTag(18);
        output.WriteMessage(AttackEnemy);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (enter_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(Enter);
      }
      if (attackEnemy_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(AttackEnemy);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(Attack other) {
      if (other == null) {
        return;
      }
      if (other.enter_ != null) {
        if (enter_ == null) {
          Enter = new global::M3Online.Domain.Enter.Enter();
        }
        Enter.MergeFrom(other.Enter);
      }
      if (other.attackEnemy_ != null) {
        if (attackEnemy_ == null) {
          AttackEnemy = new global::M3Online.Domain.Enemy.Enemy();
        }
        AttackEnemy.MergeFrom(other.AttackEnemy);
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 10: {
            if (enter_ == null) {
              Enter = new global::M3Online.Domain.Enter.Enter();
            }
            input.ReadMessage(Enter);
            break;
          }
          case 18: {
            if (attackEnemy_ == null) {
              AttackEnemy = new global::M3Online.Domain.Enemy.Enemy();
            }
            input.ReadMessage(AttackEnemy);
            break;
          }
        }
      }
    }

  }

  #endregion

}

#endregion Designer generated code
