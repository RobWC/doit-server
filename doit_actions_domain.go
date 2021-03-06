package main

//AddDomain Add Domain to datastore
func (ds *DoitServer) AddDomain(name string) (d *Domain, err error) {
	d = &Domain{Name: name}
	ds.Store.Conn.NewRecord(d)
	gormErr := ds.Store.Conn.Create(&d)
	return d, gormErr.Error
}

//AddDomainVars Add new Vars to Host
func (ds *DoitServer) AddDomainVars(id int, vars ...Var) error {
	g, err := ds.GetDomain(id)
	if err != nil {
		return err
	}
	gormErr := ds.Store.Conn.Model(&g).Association("Vars").Append(&vars)
	return gormErr.Error
}

//RemoveDomainVars Remove Vars from Host
func (ds *DoitServer) RemoveDomainVars(id int, vars ...Var) error {
	d, err := ds.GetDomain(id)
	if err != nil {
		return err
	}
	gormErr := ds.Store.Conn.Model(&d).Association("Vars").Delete(&vars)
	if gormErr != nil {
		return gormErr.Error
	}
	return nil
}

//AddDomainHosts Add new Vars to Host
func (ds *DoitServer) AddDomainHosts(id int, hosts ...Host) error {
	d, err := ds.GetDomain(id)
	if err != nil {
		return err
	}
	gormErr := ds.Store.Conn.Model(&d).Association("Hosts").Append(hosts)
	return gormErr.Error
}

//RemoveDomainHosts Remove Vars from Host
func (ds *DoitServer) RemoveDomainHosts(id int, hosts ...Host) error {
	d, err := ds.GetDomain(id)
	if err != nil {
		return err
	}
	gormErr := ds.Store.Conn.Model(&d).Association("Hosts").Delete(&hosts)
	if gormErr != nil {
		return gormErr.Error
	}
	return nil
}

//AddDomainGroups Add new Vars to Host
func (ds *DoitServer) AddDomainGroups(id int, groups ...Group) error {
	d, err := ds.GetDomain(id)
	if err != nil {
		return err
	}
	gormErr := ds.Store.Conn.Model(&d).Association("Groups").Append(groups)
	return gormErr.Error
}

//RemoveDomainGroups Remove Vars from Host
func (ds *DoitServer) RemoveDomainGroups(id int, groups ...Group) error {
	d, err := ds.GetDomain(id)
	if err != nil {
		return err
	}
	gormErr := ds.Store.Conn.Model(&d).Association("Groups").Delete(&groups)
	if gormErr != nil {
		return gormErr.Error
	}
	return nil
}

//RemoveDomain Remove Domain and its relationships to other objects
func (ds *DoitServer) RemoveDomain(d *Domain) error {
	d, err := ds.GetDomain(d.ID)
	if err != nil {
		return err
	}
	//TODO: Remove all associated objects?
	hostErr := ds.Store.Conn.Delete(&d)
	if hostErr.Error != nil {
		return hostErr.Error
	}
	return nil
}

//GetDomain Get Var from datastore
func (ds *DoitServer) GetDomain(id int) (*Domain, error) {
	d := &Domain{ID: id}
	gormErr := ds.Store.Conn.First(&d)
	if gormErr.Error != nil {
		return nil, gormErr.Error
	}
	return d, nil
}

//GetDomainByName Get Var from datastore
func (ds *DoitServer) GetDomainByName(name string) (*Domain, error) {
	d := &Domain{Name: name}
	gormErr := ds.Store.Conn.Where(&Domain{Name: name}).First(&d)
	if gormErr.Error != nil {
		return nil, gormErr.Error
	}
	return d, nil
}

//GetDomain Get Var from datastore
func (ds *DoitServer) GetDomains() ([]*Domain, error) {
	d := []*Domain{}
	gormErr := ds.Store.Conn.Find(&d)
	if gormErr.Error != nil {
		return nil, gormErr.Error
	}
	return d, nil
}
