//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdataprotection

import "encoding/json"

func unmarshalAuthCredentialsClassification(rawMsg json.RawMessage) (AuthCredentialsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b AuthCredentialsClassification
	switch m["objectType"] {
	case "SecretStoreBasedAuthCredentials":
		b = &SecretStoreBasedAuthCredentials{}
	default:
		b = &AuthCredentials{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalAuthCredentialsClassificationArray(rawMsg json.RawMessage) ([]AuthCredentialsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]AuthCredentialsClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalAuthCredentialsClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalAuthCredentialsClassificationMap(rawMsg json.RawMessage) (map[string]AuthCredentialsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages map[string]json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fMap := make(map[string]AuthCredentialsClassification, len(rawMessages))
	for key, rawMessage := range rawMessages {
		f, err := unmarshalAuthCredentialsClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fMap[key] = f
	}
	return fMap, nil
}

func unmarshalAzureBackupRecoveryPointBasedRestoreRequestClassification(rawMsg json.RawMessage) (AzureBackupRecoveryPointBasedRestoreRequestClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b AzureBackupRecoveryPointBasedRestoreRequestClassification
	switch m["objectType"] {
	case "AzureBackupRestoreWithRehydrationRequest":
		b = &AzureBackupRestoreWithRehydrationRequest{}
	default:
		b = &AzureBackupRecoveryPointBasedRestoreRequest{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalAzureBackupRecoveryPointBasedRestoreRequestClassificationArray(rawMsg json.RawMessage) ([]AzureBackupRecoveryPointBasedRestoreRequestClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]AzureBackupRecoveryPointBasedRestoreRequestClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalAzureBackupRecoveryPointBasedRestoreRequestClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalAzureBackupRecoveryPointBasedRestoreRequestClassificationMap(rawMsg json.RawMessage) (map[string]AzureBackupRecoveryPointBasedRestoreRequestClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages map[string]json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fMap := make(map[string]AzureBackupRecoveryPointBasedRestoreRequestClassification, len(rawMessages))
	for key, rawMessage := range rawMessages {
		f, err := unmarshalAzureBackupRecoveryPointBasedRestoreRequestClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fMap[key] = f
	}
	return fMap, nil
}

func unmarshalAzureBackupRecoveryPointClassification(rawMsg json.RawMessage) (AzureBackupRecoveryPointClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b AzureBackupRecoveryPointClassification
	switch m["objectType"] {
	case "AzureBackupDiscreteRecoveryPoint":
		b = &AzureBackupDiscreteRecoveryPoint{}
	default:
		b = &AzureBackupRecoveryPoint{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalAzureBackupRecoveryPointClassificationArray(rawMsg json.RawMessage) ([]AzureBackupRecoveryPointClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]AzureBackupRecoveryPointClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalAzureBackupRecoveryPointClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalAzureBackupRecoveryPointClassificationMap(rawMsg json.RawMessage) (map[string]AzureBackupRecoveryPointClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages map[string]json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fMap := make(map[string]AzureBackupRecoveryPointClassification, len(rawMessages))
	for key, rawMessage := range rawMessages {
		f, err := unmarshalAzureBackupRecoveryPointClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fMap[key] = f
	}
	return fMap, nil
}

func unmarshalAzureBackupRestoreRequestClassification(rawMsg json.RawMessage) (AzureBackupRestoreRequestClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b AzureBackupRestoreRequestClassification
	switch m["objectType"] {
	case "AzureBackupRecoveryPointBasedRestoreRequest":
		b = &AzureBackupRecoveryPointBasedRestoreRequest{}
	case "AzureBackupRecoveryTimeBasedRestoreRequest":
		b = &AzureBackupRecoveryTimeBasedRestoreRequest{}
	case "AzureBackupRestoreWithRehydrationRequest":
		b = &AzureBackupRestoreWithRehydrationRequest{}
	default:
		b = &AzureBackupRestoreRequest{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalAzureBackupRestoreRequestClassificationArray(rawMsg json.RawMessage) ([]AzureBackupRestoreRequestClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]AzureBackupRestoreRequestClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalAzureBackupRestoreRequestClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalAzureBackupRestoreRequestClassificationMap(rawMsg json.RawMessage) (map[string]AzureBackupRestoreRequestClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages map[string]json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fMap := make(map[string]AzureBackupRestoreRequestClassification, len(rawMessages))
	for key, rawMessage := range rawMessages {
		f, err := unmarshalAzureBackupRestoreRequestClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fMap[key] = f
	}
	return fMap, nil
}

func unmarshalBackupCriteriaClassification(rawMsg json.RawMessage) (BackupCriteriaClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b BackupCriteriaClassification
	switch m["objectType"] {
	case "ScheduleBasedBackupCriteria":
		b = &ScheduleBasedBackupCriteria{}
	default:
		b = &BackupCriteria{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalBackupCriteriaClassificationArray(rawMsg json.RawMessage) ([]BackupCriteriaClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]BackupCriteriaClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalBackupCriteriaClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalBackupCriteriaClassificationMap(rawMsg json.RawMessage) (map[string]BackupCriteriaClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages map[string]json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fMap := make(map[string]BackupCriteriaClassification, len(rawMessages))
	for key, rawMessage := range rawMessages {
		f, err := unmarshalBackupCriteriaClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fMap[key] = f
	}
	return fMap, nil
}

func unmarshalBackupParametersClassification(rawMsg json.RawMessage) (BackupParametersClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b BackupParametersClassification
	switch m["objectType"] {
	case "AzureBackupParams":
		b = &AzureBackupParams{}
	default:
		b = &BackupParameters{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalBackupParametersClassificationArray(rawMsg json.RawMessage) ([]BackupParametersClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]BackupParametersClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalBackupParametersClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalBackupParametersClassificationMap(rawMsg json.RawMessage) (map[string]BackupParametersClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages map[string]json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fMap := make(map[string]BackupParametersClassification, len(rawMessages))
	for key, rawMessage := range rawMessages {
		f, err := unmarshalBackupParametersClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fMap[key] = f
	}
	return fMap, nil
}

func unmarshalBaseBackupPolicyClassification(rawMsg json.RawMessage) (BaseBackupPolicyClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b BaseBackupPolicyClassification
	switch m["objectType"] {
	case "BackupPolicy":
		b = &BackupPolicy{}
	default:
		b = &BaseBackupPolicy{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalBaseBackupPolicyClassificationArray(rawMsg json.RawMessage) ([]BaseBackupPolicyClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]BaseBackupPolicyClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalBaseBackupPolicyClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalBaseBackupPolicyClassificationMap(rawMsg json.RawMessage) (map[string]BaseBackupPolicyClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages map[string]json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fMap := make(map[string]BaseBackupPolicyClassification, len(rawMessages))
	for key, rawMessage := range rawMessages {
		f, err := unmarshalBaseBackupPolicyClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fMap[key] = f
	}
	return fMap, nil
}

func unmarshalBasePolicyRuleClassification(rawMsg json.RawMessage) (BasePolicyRuleClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b BasePolicyRuleClassification
	switch m["objectType"] {
	case "AzureBackupRule":
		b = &AzureBackupRule{}
	case "AzureRetentionRule":
		b = &AzureRetentionRule{}
	default:
		b = &BasePolicyRule{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalBasePolicyRuleClassificationArray(rawMsg json.RawMessage) ([]BasePolicyRuleClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]BasePolicyRuleClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalBasePolicyRuleClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalBasePolicyRuleClassificationMap(rawMsg json.RawMessage) (map[string]BasePolicyRuleClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages map[string]json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fMap := make(map[string]BasePolicyRuleClassification, len(rawMessages))
	for key, rawMessage := range rawMessages {
		f, err := unmarshalBasePolicyRuleClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fMap[key] = f
	}
	return fMap, nil
}

func unmarshalCopyOptionClassification(rawMsg json.RawMessage) (CopyOptionClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b CopyOptionClassification
	switch m["objectType"] {
	case "CopyOnExpiryOption":
		b = &CopyOnExpiryOption{}
	case "CustomCopyOption":
		b = &CustomCopyOption{}
	case "ImmediateCopyOption":
		b = &ImmediateCopyOption{}
	default:
		b = &CopyOption{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalCopyOptionClassificationArray(rawMsg json.RawMessage) ([]CopyOptionClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]CopyOptionClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalCopyOptionClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalCopyOptionClassificationMap(rawMsg json.RawMessage) (map[string]CopyOptionClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages map[string]json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fMap := make(map[string]CopyOptionClassification, len(rawMessages))
	for key, rawMessage := range rawMessages {
		f, err := unmarshalCopyOptionClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fMap[key] = f
	}
	return fMap, nil
}

func unmarshalDataStoreParametersClassification(rawMsg json.RawMessage) (DataStoreParametersClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b DataStoreParametersClassification
	switch m["objectType"] {
	case "AzureOperationalStoreParameters":
		b = &AzureOperationalStoreParameters{}
	default:
		b = &DataStoreParameters{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalDataStoreParametersClassificationArray(rawMsg json.RawMessage) ([]DataStoreParametersClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]DataStoreParametersClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalDataStoreParametersClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalDataStoreParametersClassificationMap(rawMsg json.RawMessage) (map[string]DataStoreParametersClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages map[string]json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fMap := make(map[string]DataStoreParametersClassification, len(rawMessages))
	for key, rawMessage := range rawMessages {
		f, err := unmarshalDataStoreParametersClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fMap[key] = f
	}
	return fMap, nil
}

func unmarshalDeleteOptionClassification(rawMsg json.RawMessage) (DeleteOptionClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b DeleteOptionClassification
	switch m["objectType"] {
	case "AbsoluteDeleteOption":
		b = &AbsoluteDeleteOption{}
	default:
		b = &DeleteOption{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalDeleteOptionClassificationArray(rawMsg json.RawMessage) ([]DeleteOptionClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]DeleteOptionClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalDeleteOptionClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalDeleteOptionClassificationMap(rawMsg json.RawMessage) (map[string]DeleteOptionClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages map[string]json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fMap := make(map[string]DeleteOptionClassification, len(rawMessages))
	for key, rawMessage := range rawMessages {
		f, err := unmarshalDeleteOptionClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fMap[key] = f
	}
	return fMap, nil
}

func unmarshalFeatureValidationRequestBaseClassification(rawMsg json.RawMessage) (FeatureValidationRequestBaseClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b FeatureValidationRequestBaseClassification
	switch m["objectType"] {
	case "FeatureValidationRequest":
		b = &FeatureValidationRequest{}
	default:
		b = &FeatureValidationRequestBase{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalFeatureValidationRequestBaseClassificationArray(rawMsg json.RawMessage) ([]FeatureValidationRequestBaseClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]FeatureValidationRequestBaseClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalFeatureValidationRequestBaseClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalFeatureValidationRequestBaseClassificationMap(rawMsg json.RawMessage) (map[string]FeatureValidationRequestBaseClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages map[string]json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fMap := make(map[string]FeatureValidationRequestBaseClassification, len(rawMessages))
	for key, rawMessage := range rawMessages {
		f, err := unmarshalFeatureValidationRequestBaseClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fMap[key] = f
	}
	return fMap, nil
}

func unmarshalFeatureValidationResponseBaseClassification(rawMsg json.RawMessage) (FeatureValidationResponseBaseClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b FeatureValidationResponseBaseClassification
	switch m["objectType"] {
	case "FeatureValidationResponse":
		b = &FeatureValidationResponse{}
	default:
		b = &FeatureValidationResponseBase{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalFeatureValidationResponseBaseClassificationArray(rawMsg json.RawMessage) ([]FeatureValidationResponseBaseClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]FeatureValidationResponseBaseClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalFeatureValidationResponseBaseClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalFeatureValidationResponseBaseClassificationMap(rawMsg json.RawMessage) (map[string]FeatureValidationResponseBaseClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages map[string]json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fMap := make(map[string]FeatureValidationResponseBaseClassification, len(rawMessages))
	for key, rawMessage := range rawMessages {
		f, err := unmarshalFeatureValidationResponseBaseClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fMap[key] = f
	}
	return fMap, nil
}

func unmarshalItemLevelRestoreCriteriaClassification(rawMsg json.RawMessage) (ItemLevelRestoreCriteriaClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ItemLevelRestoreCriteriaClassification
	switch m["objectType"] {
	case "RangeBasedItemLevelRestoreCriteria":
		b = &RangeBasedItemLevelRestoreCriteria{}
	default:
		b = &ItemLevelRestoreCriteria{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalItemLevelRestoreCriteriaClassificationArray(rawMsg json.RawMessage) ([]ItemLevelRestoreCriteriaClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]ItemLevelRestoreCriteriaClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalItemLevelRestoreCriteriaClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalItemLevelRestoreCriteriaClassificationMap(rawMsg json.RawMessage) (map[string]ItemLevelRestoreCriteriaClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages map[string]json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fMap := make(map[string]ItemLevelRestoreCriteriaClassification, len(rawMessages))
	for key, rawMessage := range rawMessages {
		f, err := unmarshalItemLevelRestoreCriteriaClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fMap[key] = f
	}
	return fMap, nil
}

func unmarshalOperationExtendedInfoClassification(rawMsg json.RawMessage) (OperationExtendedInfoClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b OperationExtendedInfoClassification
	switch m["objectType"] {
	case "OperationJobExtendedInfo":
		b = &OperationJobExtendedInfo{}
	default:
		b = &OperationExtendedInfo{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalOperationExtendedInfoClassificationArray(rawMsg json.RawMessage) ([]OperationExtendedInfoClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]OperationExtendedInfoClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalOperationExtendedInfoClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalOperationExtendedInfoClassificationMap(rawMsg json.RawMessage) (map[string]OperationExtendedInfoClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages map[string]json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fMap := make(map[string]OperationExtendedInfoClassification, len(rawMessages))
	for key, rawMessage := range rawMessages {
		f, err := unmarshalOperationExtendedInfoClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fMap[key] = f
	}
	return fMap, nil
}

func unmarshalRestoreTargetInfoBaseClassification(rawMsg json.RawMessage) (RestoreTargetInfoBaseClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b RestoreTargetInfoBaseClassification
	switch m["objectType"] {
	case "ItemLevelRestoreTargetInfo":
		b = &ItemLevelRestoreTargetInfo{}
	case "RestoreFilesTargetInfo":
		b = &RestoreFilesTargetInfo{}
	case "RestoreTargetInfo":
		b = &RestoreTargetInfo{}
	default:
		b = &RestoreTargetInfoBase{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalRestoreTargetInfoBaseClassificationArray(rawMsg json.RawMessage) ([]RestoreTargetInfoBaseClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]RestoreTargetInfoBaseClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalRestoreTargetInfoBaseClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalRestoreTargetInfoBaseClassificationMap(rawMsg json.RawMessage) (map[string]RestoreTargetInfoBaseClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages map[string]json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fMap := make(map[string]RestoreTargetInfoBaseClassification, len(rawMessages))
	for key, rawMessage := range rawMessages {
		f, err := unmarshalRestoreTargetInfoBaseClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fMap[key] = f
	}
	return fMap, nil
}

func unmarshalTriggerContextClassification(rawMsg json.RawMessage) (TriggerContextClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b TriggerContextClassification
	switch m["objectType"] {
	case "AdhocBasedTriggerContext":
		b = &AdhocBasedTriggerContext{}
	case "ScheduleBasedTriggerContext":
		b = &ScheduleBasedTriggerContext{}
	default:
		b = &TriggerContext{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalTriggerContextClassificationArray(rawMsg json.RawMessage) ([]TriggerContextClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]TriggerContextClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalTriggerContextClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalTriggerContextClassificationMap(rawMsg json.RawMessage) (map[string]TriggerContextClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages map[string]json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fMap := make(map[string]TriggerContextClassification, len(rawMessages))
	for key, rawMessage := range rawMessages {
		f, err := unmarshalTriggerContextClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fMap[key] = f
	}
	return fMap, nil
}
